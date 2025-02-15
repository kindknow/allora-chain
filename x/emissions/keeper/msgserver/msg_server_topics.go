package msgserver

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	alloraMath "github.com/allora-network/allora-chain/math"
	"github.com/allora-network/allora-chain/x/emissions/metrics"
	"github.com/allora-network/allora-chain/x/emissions/types"
)

func (ms msgServer) CreateNewTopic(ctx context.Context, msg *types.CreateNewTopicRequest) (_ *types.CreateNewTopicResponse, err error) {
	defer metrics.RecordMetrics("CreateNewTopic", time.Now(), &err)

	// Validate the address
	if err := ms.k.ValidateStringIsBech32(msg.Creator); err != nil {
		return nil, err
	}
	canCreate, err := ms.k.CanCreateTopic(ctx, msg.Creator)
	if err != nil {
		return nil, err
	} else if !canCreate {
		return nil, types.ErrNotPermittedToCreateTopic
	}

	params, err := ms.k.GetParams(ctx)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "Error getting params for sender: %v", &msg.Creator)
	}
	if err := msg.Validate(params.MaxStringLength); err != nil {
		return nil, err
	}

	topicId, err := ms.k.GetNextTopicId(ctx)
	if err != nil {
		return nil, err
	}

	if msg.EpochLength < params.MinEpochLength {
		return nil, types.ErrTopicCadenceBelowMinimum
	}
	if uint64(msg.GroundTruthLag) > params.MaxUnfulfilledReputerRequests*uint64(msg.EpochLength) {
		return nil, types.ErrGroundTruthLagTooBig
	}

	// Before creating topic, transfer fee amount from creator to ecosystem bucket
	err = checkBalanceAndSendFee(ctx, ms, msg.Creator, params.CreateTopicFee)
	if err != nil {
		return nil, err
	}

	topic := types.Topic{
		Id:                       topicId,
		Creator:                  msg.Creator,
		Metadata:                 msg.Metadata,
		LossMethod:               msg.LossMethod,
		EpochLastEnded:           0,
		EpochLength:              msg.EpochLength,
		GroundTruthLag:           msg.GroundTruthLag,
		WorkerSubmissionWindow:   msg.WorkerSubmissionWindow,
		PNorm:                    msg.PNorm,
		AlphaRegret:              msg.AlphaRegret,
		AllowNegative:            msg.AllowNegative,
		Epsilon:                  msg.Epsilon,
		InitialRegret:            alloraMath.ZeroDec(),
		MeritSortitionAlpha:      msg.MeritSortitionAlpha,
		ActiveInfererQuantile:    msg.ActiveInfererQuantile,
		ActiveForecasterQuantile: msg.ActiveForecasterQuantile,
		ActiveReputerQuantile:    msg.ActiveReputerQuantile,
	}
	_, err = ms.k.IncrementTopicId(ctx)
	if err != nil {
		return nil, err
	}
	if err := ms.k.SetTopic(ctx, topicId, topic); err != nil {
		return nil, err
	}

	// Turn topic whitelist on by default so no one can squeeze in payloads before an admin notices or can act
	if msg.EnableWorkerWhitelist {
		err = ms.k.EnableTopicWorkerWhitelist(ctx, topicId)
		if err != nil {
			return nil, err
		}
	}
	if msg.EnableReputerWhitelist {
		err = ms.k.EnableTopicReputerWhitelist(ctx, topicId)
		if err != nil {
			return nil, err
		}
	}

	err = ms.k.AddTopicFeeRevenue(ctx, topicId, params.CreateTopicFee)
	return &types.CreateNewTopicResponse{TopicId: topicId}, err
}
