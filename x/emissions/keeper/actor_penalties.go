package keeper

import (
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
		ctx,
		CountWorkerContiguousMissedEpochs,
		func(topicId TopicId) (alloraMath.Dec, error) {
			return k.GetTopicInitialInfererEmaScore(ctx, topicId)
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
		ctx,
		CountWorkerContiguousMissedEpochs,
		func(topicId TopicId) (alloraMath.Dec, error) {
			return k.GetTopicInitialForecasterEmaScore(ctx, topicId)
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
		ctx,
		CountReputerContiguousMissedEpochs,
		func(topicId TopicId) (alloraMath.Dec, error) {
			return k.GetTopicInitialReputerEmaScore(ctx, topicId)
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
	ctx sdk.Context,
	missedEpochsFn func(topic types.Topic, lastSubmittedNonce int64) int64,
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
	if err != nil {
		return types.Score{}, err
	}
	emaScore.BlockHeight = block

	beforePenalty := emaScore
	emaScore.Score, err = applyPenalty(topic, penalty, emaScore.Score, missedEpochs)
	if err != nil {
		return types.Score{}, err
	}

	ctx.Logger().Debug("apply liveness penalty on actor",
		"nonce", block,
		"penalty", penalty,
		"before", beforePenalty,
		"after", emaScore,
	)

	// Save the penalised EMA score
	return emaScore, setScoreFn(topic.Id, emaScore)
}

// applyPenalty applies the penalty to the EMA score for the given number of missed epochs while staying above provided limit.
func applyPenalty(topic types.Topic, penalty, emaScore alloraMath.Dec, missedEpochs int64) (alloraMath.Dec, error) {
	return alloraMath.NCalcEma(topic.MeritSortitionAlpha, penalty, emaScore, uint64(missedEpochs))
}

// CountWorkerContiguousMissedEpochs counts the number of contiguous missed epochs prior to the given nonce, given the
// actor last submission.
func CountWorkerContiguousMissedEpochs(topic types.Topic, lastSubmittedNonce int64) int64 {
	prevEpochStart := topic.EpochLastEnded - topic.EpochLength
	return countContiguousMissedEpochs(prevEpochStart, topic.EpochLength, lastSubmittedNonce)
}

// CountReputerContiguousMissedEpochs counts the number of contiguous missed epochs prior to the given nonce, given the
// actor last submission.
func CountReputerContiguousMissedEpochs(topic types.Topic, lastSubmittedNonce int64) int64 {
	prevEpochStart := topic.EpochLastEnded - topic.EpochLength - topic.GroundTruthLag
	return countContiguousMissedEpochs(prevEpochStart, topic.EpochLength, lastSubmittedNonce)
}

func countContiguousMissedEpochs(prevEpochStart, epochLength, lastSubmittedNonce int64) int64 {
	if lastSubmittedNonce >= prevEpochStart {
		return 0
	}

	return (prevEpochStart-1-lastSubmittedNonce)/epochLength + 1
}
