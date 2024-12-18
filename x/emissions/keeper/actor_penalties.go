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
	return MayPenaliseActor(
		CountWorkerContiguousMissedEpochs,
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
	return MayPenaliseActor(
		CountWorkerContiguousMissedEpochs,
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

// MayPenaliseReputer penalises a reputer for missing previous epochs. It saves and returns the new EMA score.
// If the reputer didn't miss any epochs this is a no-op, the EMA score is returned as is.
func (k *Keeper) MayPenaliseReputer(
	ctx sdk.Context,
	topic types.Topic,
	block types.BlockHeight,
	emaScore types.Score,
) (types.Score, error) {
	return MayPenaliseActor(
		CountReputerContiguousMissedEpochs,
		func(topicId TopicId) (alloraMath.Dec, error) {
			return k.initialReputerEmaScore.Get(ctx, topicId)
		},
		func(topicId TopicId, score types.Score) error {
			return k.SetReputerScoreEma(ctx, topicId, score.Address, score)
		},
		topic,
		block,
		emaScore,
	)
}

func MayPenaliseActor(
	missedEpochsFn func(topic types.Topic, lastSubmissionHeight int64) int64,
	getPenaltyFn func(topicId TopicId) (alloraMath.Dec, error),
	setScoreFn func(topicId TopicId, score types.Score) error,
	topic types.Topic,
	block types.BlockHeight,
	emaScore types.Score,
) (types.Score, error) {
	missedEpochs := missedEpochsFn(topic, emaScore.BlockHeight)
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
	emaScore.Score, err = applyPenalty(topic, penalty, emaScore.Score, missedEpochs)
	if err != nil {
		return types.Score{}, err
	}

	// Save the penalised EMA score
	return emaScore, setScoreFn(topic.Id, emaScore)
}

// applyPenalty applies the penalty to the EMA score for the given number of missed epochs while staying above provided limit.
func applyPenalty(topic types.Topic, penalty, emaScore alloraMath.Dec, missedEpochs int64) (alloraMath.Dec, error) {
	return alloraMath.NCalcEma(topic.MeritSortitionAlpha, penalty, emaScore, uint64(missedEpochs))
}

// CountWorkerContiguousMissedEpochs counts the number of contiguous missed epochs prior to the given nonce, given the
// actor last submission.
func CountWorkerContiguousMissedEpochs(topic types.Topic, lastSubmissionHeight int64) int64 {
	prevEpochStart := topic.EpochLastEnded - topic.EpochLength
	return countContiguousMissedEpochs(prevEpochStart, topic.EpochLength, lastSubmissionHeight)
}

// CountReputerContiguousMissedEpochs counts the number of contiguous missed epochs prior to the given nonce, given the
// actor last submission.
func CountReputerContiguousMissedEpochs(topic types.Topic, lastSubmissionHeight int64) int64 {
	prevEpochStart := topic.EpochLastEnded - topic.EpochLength + topic.WorkerSubmissionWindow + topic.GroundTruthLag
	return countContiguousMissedEpochs(prevEpochStart, topic.EpochLength, lastSubmissionHeight)
}

func countContiguousMissedEpochs(prevEpochStart, epochLength, lastSubmissionHeight int64) int64 {
	if lastSubmissionHeight >= prevEpochStart {
		return 0
	}

	return (prevEpochStart-1-lastSubmissionHeight)/epochLength + 1
}
