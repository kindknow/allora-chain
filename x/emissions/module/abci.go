package module

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/errors"
	allorautils "github.com/allora-network/allora-chain/x/emissions/keeper/actor_utils"
	"github.com/allora-network/allora-chain/x/emissions/module/rewards"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx context.Context, am AppModule) error {
	defer telemetry.ModuleMeasureSince(emissionstypes.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	sdkCtx.Logger().Debug(
		fmt.Sprintf("\n ---------------- Emissions EndBlock %d ------------------- \n",
			blockHeight))
	moduleParams, err := am.keeper.GetParams(sdkCtx)
	if err != nil {
		sdkCtx.Logger().Error("Error Getting module params", err)
		return err
	}
	// Remove Stakers that have been wanting to unstake this block. They no longer get paid rewards
	err = RemoveStakes(sdkCtx, blockHeight, am.keeper, moduleParams.HalfMaxProcessStakeRemovalsEndBlock)
	if err != nil {
		sdkCtx.Logger().Error("Error removing stakes: ", err)
	}
	err = RemoveDelegateStakes(sdkCtx, blockHeight, am.keeper, moduleParams.HalfMaxProcessStakeRemovalsEndBlock)
	if err != nil {
		sdkCtx.Logger().Error("Error removing delegate stakes: ", err)
	}

	// Get unnormalized weights of active topics and the sum weight and revenue they have generated
	weights, sumWeight, totalRevenue, err := rewards.GetAndUpdateActiveTopicWeights(sdkCtx, am.keeper, blockHeight)
	if err != nil {
		return errors.Wrapf(err, "Weights error")
	}

	sdkCtx.Logger().Debug(fmt.Sprintf("ABCI EndBlocker %d: Total Revenue: %v, Sum Weight: %v", blockHeight, totalRevenue, sumWeight))

	err = rewards.UpdateNoncesOfActiveTopics(
		sdkCtx,
		am.keeper,
		blockHeight,
		weights,
	)
	if err != nil {
		sdkCtx.Logger().Error("Error applying function on all rewardable topics: ", err)
		return err
	}

	// REWARDS (will internally filter any non-RewardReady topics)
	err = rewards.EmitRewards(rewards.EmitRewardsArgs{
		Ctx:          sdkCtx,
		K:            am.keeper,
		ModuleParams: moduleParams,
		BlockHeight:  blockHeight,
		Weights:      weights,
		SumWeight:    sumWeight,
		TotalRevenue: totalRevenue,
	})
	if err != nil {
		sdkCtx.Logger().Error("Error calculating global emission per topic: ", err)
		return errors.Wrapf(err, "Rewards error")
	}
	// Close any open windows due this blockHeight
	workerWindowsToClose := am.keeper.GetWorkerWindowTopicIds(sdkCtx, blockHeight)
	if len(workerWindowsToClose.TopicIds) > 0 {
		for _, topicId := range workerWindowsToClose.TopicIds {
			sdkCtx.Logger().Info(fmt.Sprintf("ABCI EndBlocker: Worker close cadence met for topic: %d", topicId))
			// Check if there is an unfulfilled nonce
			nonces, err := am.keeper.GetUnfulfilledWorkerNonces(sdkCtx, topicId)
			if err != nil {
				sdkCtx.Logger().Warn(fmt.Sprintf("Error getting unfulfilled worker nonces: %s", err.Error()))
				continue
			} else if len(nonces.Nonces) == 0 {
				// No nonces to fulfill
				continue
			} else {
				topic, err := am.keeper.GetTopic(sdkCtx, topicId)
				if err != nil {
					sdkCtx.Logger().Warn(fmt.Sprintf("Error getting topic: %s", err.Error()))
					continue
				}
				for _, nonce := range nonces.Nonces {
					// Skip rest of logic if the worker submission window is still open (i.e. don't close the window yet)
					// Note: worker window exclusive of the end block height so we ensure the window is closed at this point.
					if am.keeper.BlockWithinWorkerSubmissionWindowOfNonce(topic, *nonce, blockHeight) {
						sdkCtx.Logger().Debug(fmt.Sprintf("ABCI EndBlocker %d: Worker window still open for topic: %d, nonce: %v", blockHeight, topicId, nonce))
						continue
					}
					sdkCtx.Logger().Debug(fmt.Sprintf("ABCI EndBlocker %d: Closing Worker window for topic: %d, nonce: %v", blockHeight, topicId, nonce))
					err := allorautils.CloseWorkerNonce(&am.keeper, sdkCtx, topic, *nonce)
					if err != nil {
						sdkCtx.Logger().Info(fmt.Sprintf("Error closing worker nonce, proactively fulfilling: %s", err.Error()))
						// Proactively close the nonce
						fulfilledNonce, err := am.keeper.FulfillWorkerNonce(sdkCtx, topicId, nonce)
						if err != nil {
							sdkCtx.Logger().Warn(fmt.Sprintf("Error fulfilling worker nonce: %s", err.Error()))
						} else {
							sdkCtx.Logger().Debug(fmt.Sprintf("Fulfilled: %t, worker nonce: %v", fulfilledNonce, nonce))
						}
					}
				}
			}
		}
		err = am.keeper.DeleteWorkerWindowBlockHeight(sdkCtx, blockHeight)
		if err != nil {
			sdkCtx.Logger().Warn(fmt.Sprintf("Error deleting worker window blockheight: %s", err.Error()))
		}
	}
	return nil
}
