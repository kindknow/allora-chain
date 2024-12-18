package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	alloraMath "github.com/allora-network/allora-chain/math"
	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MayPenaliseInferer penalises an inferer for missing previous epochs. It saves and returns the new EMA score.
// If the inferer didn't miss any epochs this is a no-op, the EMA score is returned as is.
func (k *Keeper) MayPenaliseInferer(
	ctx sdk.Context,
	topic types.Topic,
	block types.BlockHeight,
	emaScore types.Score,
) (types.Score, error) {
	return mayPenaliseWorker(
		func(topicId TopicId) (alloraMath.Dec, error) {
			return k.initialInfererEmaScore.Get(ctx, topicId)
		},
		func(topicId TopicId, score types.Score) error {
			return k.SetInfererScoreEma(ctx, topicId, score.Address, score)
		},
		topic,
		block,
		emaScore,
	)
}

// MayPenaliseForecaster penalises a forecaster for missing previous epochs. It saves and returns the new EMA score.
// If the forecaster didn't miss any epochs this is a no-op, the EMA score is returned as is.
func (k *Keeper) MayPenaliseForecaster(
	ctx sdk.Context,
	topic types.Topic,
	block types.BlockHeight,
	emaScore types.Score,
) (types.Score, error) {
	return mayPenaliseWorker(
		func(topicId TopicId) (alloraMath.Dec, error) {
			return k.initialForecasterEmaScore.Get(ctx, topicId)
		},
		func(topicId TopicId, score types.Score) error {
			return k.SetForecasterScoreEma(ctx, topicId, score.Address, score)
		},
		topic,
		block,
		emaScore,
	)
}

func mayPenaliseWorker(
	getPenaltyFn func(topicId TopicId) (alloraMath.Dec, error),
	setScoreFn func(topicId TopicId, score types.Score) error,
	topic types.Topic,
	block types.BlockHeight,
	emaScore types.Score,
) (types.Score, error) {
	missedEpochs := countWorkerContiguousMissedEpochs(topic, emaScore.BlockHeight)
	// No missed epochs == no penalty
	if missedEpochs == 0 {
		return emaScore, nil
	}

	penalty, err := getPenaltyFn(topic.Id)
	// No penalty set, no penalty
	if errors.Is(err, collections.ErrNotFound) {
		return emaScore, nil
	}
	if err != nil {
		return types.Score{}, err
	}
	emaScore.BlockHeight = block
	// TODO: Provide the limit as the initial score
	emaScore.Score, err = applyPenalty(topic, penalty, alloraMath.ZeroDec(), emaScore.Score, missedEpochs)
	if err != nil {
		return types.Score{}, err
	}

	// Save the penalised EMA score
	return emaScore, setScoreFn(topic.Id, emaScore)
}

// applyPenalty applies the penalty to the EMA score for the given number of missed epochs while staying above provided limit.
func applyPenalty(topic types.Topic, penalty, limit, emaScore alloraMath.Dec, missedEpochs int64) (alloraMath.Dec, error) {
	for i := int64(0); i < missedEpochs; i++ {
		penalisedScore, err := alloraMath.CalcEma(topic.MeritSortitionAlpha, penalty, emaScore, false)
		if err != nil {
			return alloraMath.ZeroDec(), err
		}

		if penalisedScore.Lt(limit) {
			break
		}
		emaScore = penalisedScore
	}

	return emaScore, nil
}

// CountWorkerContiguousMissedEpochs counts the number of contiguous missed epochs prior to the given nonce, given the
// last worker submission and the current block heights.
func countWorkerContiguousMissedEpochs(topic types.Topic, lastSubmissionHeight int64) int64 {
	prevEpochStart := topic.EpochLastEnded - topic.EpochLength
	if lastSubmissionHeight >= prevEpochStart {
		return 0
	}

	return (prevEpochStart - lastSubmissionHeight) / topic.EpochLength
}
