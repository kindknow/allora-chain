package msgserver

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	actorutils "github.com/allora-network/allora-chain/x/emissions/keeper/actor_utils"
	"github.com/allora-network/allora-chain/x/emissions/metrics"
	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// A tx function that accepts a individual inference and forecast and possibly returns an error
// Need to call this once per forecaster per topic inference solicitation round because protobuf does not nested repeated fields
// Only 1 payload per registered worker is kept, ignore the rest. In particular, take the first payload from each
// registered worker and none from any unregistered actor.
// Signatures, anti-sybil procedures, and "skimming of only the top few workers by EMA score descending" should be done here.
func (ms msgServer) InsertWorkerPayload(ctx context.Context, msg *types.InsertWorkerPayloadRequest) (_ *types.InsertWorkerPayloadResponse, err error) {
	defer metrics.RecordMetrics("InsertWorkerPayload", time.Now(), &err)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "Error validating sender address")
	}
	err = ms.k.ValidateStringIsBech32(msg.WorkerDataBundle.Worker)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "Error validating worker address")
	}
	canSubmit, err := ms.k.CanSubmitWorkerPayload(ctx, msg.WorkerDataBundle.TopicId, msg.WorkerDataBundle.Worker)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "Error checking if worker can submit payload")
	} else if !canSubmit {
		return nil, errorsmod.Wrapf(types.ErrNotPermittedToSubmitWorkerPayload, "Worker is not permitted to submit payload")
	}

	blockHeight := sdkCtx.BlockHeight()
	err = msg.WorkerDataBundle.Validate()
	if err != nil {
		return nil, errorsmod.Wrapf(err,
			"Worker invalid data for block: %d", blockHeight)
	}

	moduleParams, err := ms.k.GetParams(ctx)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "Error getting params")
	}
	err = checkInputLength(moduleParams.MaxSerializedMsgLength, msg)
	if err != nil {
		return nil, err
	}

	nonce := msg.WorkerDataBundle.Nonce
	topicId := msg.WorkerDataBundle.TopicId

	topic, err := ms.k.GetTopic(ctx, topicId)
	if err != nil {
		return nil, types.ErrInvalidTopicId
	}
	nonceUnfulfilled, err := ms.k.IsWorkerNonceUnfulfilled(ctx, topicId, nonce)
	if err != nil {
		return nil, err
	} else if !nonceUnfulfilled {
		return nil, types.ErrUnfulfilledNonceNotFound
	}

	// Note: this is exclusive of the end block height
	if !ms.k.BlockWithinWorkerSubmissionWindowOfNonce(topic, *nonce, blockHeight) {
		return nil, errorsmod.Wrapf(
			types.ErrWorkerNonceWindowNotAvailable,
			"Worker window not open for topic: %d, current block %d, start window: %d, end window: %d",
			topicId, blockHeight, nonce.BlockHeight, nonce.BlockHeight+topic.WorkerSubmissionWindow,
		)
	}

	isWorkerRegistered, err := ms.k.IsWorkerRegisteredInTopic(ctx, topicId, msg.WorkerDataBundle.Worker)
	if err != nil {
		return nil, err
	} else if !isWorkerRegistered {
		return nil, errorsmod.Wrapf(types.ErrAddressNotRegistered, "worker is not registered in this topic")
	}

	err = sendEffectiveRevenueActivateTopicIfWeightSufficient(ctx, ms, msg.Sender, topicId, moduleParams.DataSendingFee)
	if err != nil {
		return nil, err
	}

	// Process Inferences
	if msg.WorkerDataBundle.InferenceForecastsBundle.Inference != nil {
		inference := msg.WorkerDataBundle.InferenceForecastsBundle.Inference
		if inference == nil {
			return nil, errorsmod.Wrapf(types.ErrNoValidInferences, "Inference not found")
		}
		if inference.TopicId != msg.WorkerDataBundle.TopicId {
			return nil, errorsmod.Wrapf(types.ErrInvalidTopicId,
				"inferer not using the same topic as bundle")
		}

		err = ms.k.AppendInference(sdkCtx, topic, nonce.BlockHeight, inference, moduleParams.MaxTopInferersToReward)
		if err != nil {
			return nil, errorsmod.Wrapf(err, "Error appending inference")
		}
	}

	// Process Forecasts
	if msg.WorkerDataBundle.InferenceForecastsBundle.Forecast != nil {
		forecast := msg.WorkerDataBundle.InferenceForecastsBundle.Forecast
		if len(forecast.ForecastElements) == 0 {
			return nil, errorsmod.Wrapf(types.ErrNoValidForecastElements, "No valid forecast elements found in Forecast")
		}
		if forecast.TopicId != msg.WorkerDataBundle.TopicId {
			return nil, errorsmod.Wrapf(types.ErrInvalidTopicId, "forecaster not using the same topic as bundle")
		}

		// Limit forecast elements to top inferers
		latestScoresForForecastedInferers := make([]types.Score, 0)
		for _, el := range forecast.ForecastElements {
			score, err := ms.k.GetInfererScoreEma(ctx, forecast.TopicId, el.Inferer)
			if err != nil {
				continue
			}
			latestScoresForForecastedInferers = append(latestScoresForForecastedInferers, score)
		}

		_, _, topNInferer := actorutils.FindTopNByScoreDesc(
			sdkCtx,
			moduleParams.MaxElementsPerForecast,
			latestScoresForForecastedInferers,
			forecast.BlockHeight,
		)

		// Remove duplicate forecast elements
		acceptedForecastElements := make([]*types.ForecastElement, 0)
		seenInferers := make(map[string]bool)
		for _, el := range forecast.ForecastElements {
			// Check if the forecasted inferer is registered in the topic
			isInfererRegistered, err := ms.k.IsWorkerRegisteredInTopic(ctx, topicId, el.Inferer)
			if err != nil {
				return nil, err
			}
			if !isInfererRegistered {
				return nil, errorsmod.Wrapf(err,
					"Error forecasted inferer address is not registered in this topic")
			}

			notAlreadySeen := !seenInferers[el.Inferer]
			_, isTopInferer := topNInferer[el.Inferer]
			if notAlreadySeen && isTopInferer {
				acceptedForecastElements = append(acceptedForecastElements, el)
				seenInferers[el.Inferer] = true
			}
		}

		if len(acceptedForecastElements) > 0 {
			forecast.ForecastElements = acceptedForecastElements
			err = ms.k.AppendForecast(sdkCtx, topic, nonce.BlockHeight, forecast, moduleParams.MaxTopForecastersToReward)
			if err != nil {
				return nil, errorsmod.Wrapf(err,
					"Error appending forecast")
			}
		}
	}
	return &types.InsertWorkerPayloadResponse{}, nil
}
