package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	statev7 "github.com/allora-network/allora-chain/x/emissions/api/emissions/v7"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: statev7.QueryService_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetParams",
					Use:       "params",
					Short:     "Get the current module parameters",
				},
				{
					RpcMethod: "GetNextTopicId",
					Use:       "next-topic-id",
					Short:     "Get next topic id",
				},
				{
					RpcMethod: "GetTopic",
					Use:       "topic [topic_id]",
					Short:     "Get topic by topic id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "TopicExists",
					Use:       "topic-exists [topic_id]",
					Short:     "True if topic exists at given id, else false",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "IsTopicActive",
					Use:       "is-topic-active [topic_id]",
					Short:     "True if topic is active, else false",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetReputerStakeInTopic",
					Use:       "stake-in-topic-reputer [address] [topic_id]",
					Short:     "Get reputer stake in a topic, including stake delegated to them in that topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetDelegateStakeInTopicInReputer",
					Use:       "stake-total-delegated-in-topic-reputer [reputer_address] [topic_id]",
					Short:     "Get total delegate stake in a topic and reputer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "reputer_address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetDelegateRewardPerShare",
					Use:       "delegate-reward-per-share [topic_id] [reputer_address]",
					Short:     "Get total delegate reward per share stake in a reputer for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "GetDelegateStakePlacement",
					Use:       "delegate-stake-placement [topic_id] [delegator] [target]",
					Short:     "Get amount of token delegated to a target by a delegator in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "delegator"},
						{ProtoField: "target"},
					},
				},
				{
					RpcMethod: "GetDelegateStakeRemoval",
					Use:       "delegate-stake-removal [block_height] [topic_id] [delegator] [reputer]",
					Short:     "Get current state of a pending delegate stake removal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "block_height"},
						{ProtoField: "topic_id"},
						{ProtoField: "delegator"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "GetDelegateStakeUponReputer",
					Use:       "delegate-stake-on-reputer [topic_id] [target]",
					Short:     "Get total amount of token delegated to a target reputer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "target"},
					},
				},
				{
					RpcMethod: "GetForecastScoresUntilBlock",
					Use:       "forecast-scores-until-block [topic_id] [block_height]",
					Short:     "Get all saved scores for all forecasters for a topic descending until a given past block height. Number of forecasts is limited by MaxSamplesToScaleScores",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetForecasterNetworkRegret",
					Use:       "forecaster-regret [topic_id] [worker]",
					Short:     "Get current network regret for given forecaster",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "worker"},
					},
				},
				{
					RpcMethod: "GetInferenceScoresUntilBlock",
					Use:       "inference-scores-until-block [topic_id] [block_height]",
					Short:     "Get all saved scores for all inferers for a topic descending until a given past block height. Number of forecasts is limited by MaxSamplesToScaleScores",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetInfererNetworkRegret",
					Use:       "inferer-regret [topic_id] [actor_id]",
					Short:     "Get current network regret for given inferer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "actor_id"},
					},
				},
				{
					RpcMethod: "IsReputerNonceUnfulfilled",
					Use:       "reputer-nonce-unfulfilled [topic_id] [block_height]",
					Short:     "True if reputer nonce is unfulfilled (still awaiting a reputer response), else false",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "IsWorkerNonceUnfulfilled",
					Use:       "worker-nonce-unfulfilled [topic_id] [block_height]",
					Short:     "True if worker nonce is unfulfilled (still awaiting a worker response), else false",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetLatestAvailableNetworkInferences",
					Use:       "latest-available-network-inferences [topic_id]",
					Short:     "Returns network inference only if all available information to compute the inference is present",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetLatestAvailableNetworkInferencesOutlierResistant",
					Use:       "latest-available-network-inferences-outlier-resistant [topic_id]",
					Short:     "Returns network inference only if all available information to compute the inference is present, removing outliers",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetForecasterScoreEma",
					Use:       "forecaster-score-ema [topic_id] [forecaster]",
					Short:     "Returns latest score for a forecaster in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "forecaster"},
					},
				},
				{
					RpcMethod: "GetPreviousTopicQuantileForecasterScoreEma",
					Use:       "topic-quantile-forecaster-score [topic_id]",
					Short:     "Returns topic-quantile score ema among the previous top forecasters by score EMA",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetCurrentLowestForecasterScore",
					Use:       "current-lowest-forecaster-score [topic_id]",
					Short:     "Returns current lowest score for a forecaster in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetInfererScoreEma",
					Use:       "inferer-score-ema [topic_id] [inferer]",
					Short:     "Returns latest score for a inferer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "inferer"},
					},
				},
				{
					RpcMethod: "GetPreviousTopicQuantileInfererScoreEma",
					Use:       "topic-quantile-inferer-score [topic_id]",
					Short:     "Returns topic-quantile score ema among the previous top inferers by score EMA",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetCurrentLowestInfererScore",
					Use:       "current-lowest-inferer-score [topic_id]",
					Short:     "Returns current lowest score for a inferer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetReputerScoreEma",
					Use:       "reputer-score-ema [topic_id] [reputer]",
					Short:     "Returns latest score for a reputer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "GetPreviousTopicQuantileReputerScoreEma",
					Use:       "topic-quantile-reputer-score [topic_id]",
					Short:     "Returns topic-quantile score ema among the previous top reputers by score EMA",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetCurrentLowestReputerScore",
					Use:       "current-lowest-reputer-score [topic_id]",
					Short:     "Returns current lowest score for a reputer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetLatestTopicInferences",
					Use:       "latest-topic-raw-inferences [topic_id]",
					Short:     "Returns latest round of raw inferences from workers topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetListeningCoefficient",
					Use:       "listening-coefficient [topic_id] [reputer]",
					Short:     "Returns current listening coefficient for a given reputer. Default to 1 if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "GetMultiReputerStakeInTopic",
					Use:       "multi-reputer-stake [addresses] [topic_id]",
					Short:     "Returns stakes for each reputer in a given list. List can be up to MaxPageLimit in length. Default to 0 if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "addresses"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetOneInForecasterNetworkRegret",
					Use:       "one-in-forecaster-regret [topic_id] [forecaster] [inferer]",
					Short:     "Returns regret born from including [forecaster]'s implied inference in a batch with [inferer]. Default to topic InitialRegret if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "forecaster"},
						{ProtoField: "inferer"},
					},
				},
				{
					RpcMethod: "GetNaiveInfererNetworkRegret",
					Use:       "naive-inferer-network-regret [topic_id] [inferer]",
					Short:     "Returns regret born from including [inferer]'s naive inference in a batch. Default to topic InitialRegret if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "inferer"},
					},
				},
				{
					RpcMethod: "GetOneOutInfererInfererNetworkRegret",
					Use:       "one-out-inferer-inferer-network-regret [topic_id] [one_out_inferer] [inferer]",
					Short:     "Returns regret born from including [one_out_inferer]'s implied inference in a batch with [inferer]. Default to topic InitialRegret if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "one_out_inferer"},
						{ProtoField: "inferer"},
					},
				},
				{
					RpcMethod: "GetOneOutInfererForecasterNetworkRegret",
					Use:       "one-out-inferer-forecaster-network-regret [topic_id] [one_out_inferer] [forecaster]",
					Short:     "Returns regret born from including [one_out_inferer]'s implied inference in a batch with [forecaster]. Default to topic InitialRegret if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "one_out_inferer"},
						{ProtoField: "forecaster"},
					},
				},
				{
					RpcMethod: "GetOneOutForecasterInfererNetworkRegret",
					Use:       "one-out-forecaster-inferer-network-regret [topic_id] [one_out_forecaster] [inferer]",
					Short:     "Returns regret born from including [one_out_forecaster]'s implied inference in a batch with [inferer]. Default to topic InitialRegret if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "one_out_forecaster"},
						{ProtoField: "inferer"},
					},
				},
				{
					RpcMethod: "GetOneOutForecasterForecasterNetworkRegret",
					Use:       "one-out-forecaster-forecaster-network-regret [topic_id] [one_out_forecaster] [forecaster]",
					Short:     "Returns regret born from including [one_out_forecaster]'s implied inference in a batch with [forecaster]. Default to topic InitialRegret if does not exist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "one_out_forecaster"},
						{ProtoField: "forecaster"},
					},
				},
				{
					RpcMethod: "GetPreviousForecastRewardFraction",
					Use:       "previous-forecaster-reward-fraction [topic_id] [worker]",
					Short:     "Return previous reward fraction for actor",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "worker"},
					},
				},
				{
					RpcMethod: "GetPreviousInferenceRewardFraction",
					Use:       "previous-inference-reward-fraction [topic_id] [worker]",
					Short:     "Return previous reward fraction for actor",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "worker"},
					},
				},
				{
					RpcMethod: "GetPreviousPercentageRewardToStakedReputers",
					Use:       "previous-percentage-reputer-reward",
					Short:     "Return previous percent reward paid to staked reputers",
				},
				{
					RpcMethod: "GetPreviousReputerRewardFraction",
					Use:       "previous-reputer-reward-fraction [topic_id] [reputer]",
					Short:     "Return previous reward fraction for actor",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "GetPreviousTopicWeight",
					Use:       "previous-topic-weight [topic_id]",
					Short:     "Return previous topic weight. Useful for extrapolating future and previous topic weight and the topic's likelihood for churn",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod:      "GetTotalSumPreviousTopicWeights",
					Use:            "previous-total-topic-weights",
					Short:          "Return previous total topic weights sum.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod: "GetReputerLossBundlesAtBlock",
					Use:       "reputer-loss-bundle [topic_id] [block_height]",
					Short:     "Return reputer loss bundle at block height. May not exist if it was already pruned",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetReputersScoresAtBlock",
					Use:       "reputer-scores-at-block [topic_id] [block_height]",
					Short:     "Return reputer scores at block. Note: the chain only stores up to MaxSamplesToScaleScores many scores per actor type per topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetStakeRemovalForReputerAndTopicId",
					Use:       "stake-removal [reputer] [topic_id]",
					Short:     "Return stake removal information for reputer in topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "reputer"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetStakeReputerAuthority",
					Use:       "reputer-authority [topic_id] [reputer]",
					Short:     "Return total stake on reputer in a topic, including delegate stake and their own",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "GetTopicFeeRevenue",
					Use:       "topic-fee-revenue [topic_id]",
					Short:     "Return effective fee revenue for a topic i.e. the total fees collected by the topic less an exponential decay of the fees over time. This is the impact of topic fees on the topic's weight",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetTopicRewardNonce",
					Use:       "topic-reward-nonce [topic_id]",
					Short:     "If a topic is rewardable, then this is the nonce that will be used to calculate topic rewards. The actors that participated in the worker/reputer rounds started at this nonce (block height) will be rewarded",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetTopicStake",
					Use:       "topic-stake [topic_id]",
					Short:     "Return total stake in topic including delegate stake",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetTotalRewardToDistribute",
					Use:       "total-rewards",
					Short:     "Return total rewards to be distributed among all rewardable topics in the block",
				},
				{
					RpcMethod: "GetUnfulfilledReputerNonces",
					Use:       "unfulfilled-reputer-nonces [topic_id]",
					Short:     "Return topic reputer nonces that have yet to be fulfilled",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetUnfulfilledWorkerNonces",
					Use:       "unfulfilled-worker-nonces [topic_id]",
					Short:     "Return topic worker nonces that have yet to be fulfilled",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetWorkerForecastScoresAtBlock",
					Use:       "forecast-scores [topic_id] [block_height]",
					Short:     "Return scores for topic worker at a block height. Default is empty. May not exist if it was already pruned",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetWorkerInferenceScoresAtBlock",
					Use:       "inference-scores [topic_id] [block_height]",
					Short:     "Return scores for topic worker at a block height. Default is empty. May not exist if it was already pruned",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetStakeFromReputerInTopicInSelf",
					Use:       "stake-reputer-in-topic-self [reputer_address] [topic_id]",
					Short:     "Get the stake of a reputer in a topic that they put on themselves",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "reputer_address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetStakeFromDelegatorInTopicInReputer",
					Use:       "stake-delegator-in-topic-reputer [delegator_address] [reputer_address] [topic_id]",
					Short:     "Get amount of stake from delegator in a topic for a reputer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "delegator_address"},
						{ProtoField: "reputer_address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetStakeFromDelegatorInTopic",
					Use:       "stake-delegator-in-topic [delegator_address] [topic_id]",
					Short:     "Get amount of stake in a topic for a delegator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "delegator_address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetWorkerLatestInferenceByTopicId",
					Use:       "worker-latest-inference [topic_id] [worker_address]",
					Short:     "Get the latest inference for a given worker and topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "worker_address"},
					},
				},
				{
					RpcMethod: "GetInferencesAtBlock",
					Use:       "inferences-at-block [topic_id] [block_height]",
					Short:     "Get All Inferences produced for a topic in a particular timestamp",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetWorkerNodeInfo",
					Use:       "worker-info [address]",
					Short:     "Get node info for worker node",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "GetReputerNodeInfo",
					Use:       "reputer-info [address]",
					Short:     "Get node info for reputer node",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsWorkerRegisteredInTopicId",
					Use:       "is-worker-registered [topic_id] [address]",
					Short:     "True if worker is registered in the topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsReputerRegisteredInTopicId",
					Use:       "is-reputer-registered [topic_id] [address]",
					Short:     "True if reputer is registered in the topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "GetTotalStake",
					Use:       "total-stake",
					Short:     "Get the total amount of staked tokens by all participants in the network",
				},
				{
					RpcMethod: "GetForecastsAtBlock",
					Use:       "forecasts-at-block [topic_id] [block]",
					Short:     "Get the Forecasts for a topic at block height ",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetNetworkInferencesAtBlock",
					Use:       "network-inferences-at-block [topic_id] [block_height_last_inference]",
					Short:     "Get the Network Inferences for a topic at a block height where the last inference was made",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height_last_inference"},
					},
				},
				{
					RpcMethod: "GetNetworkInferencesAtBlockOutlierResistant",
					Use:       "network-inferences-at-block-outlier-resistant [topic_id] [block_height_last_inference]",
					Short:     "Get the Network Inferences for a topic at a block height where the last inference was made removing outliers",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height_last_inference"},
					},
				},
				{
					RpcMethod: "GetLatestNetworkInferences",
					Use:       "latest-network-inferences [topic_id]",
					Short:     "Get the latest Network inferences and weights for a topic. Will return whatever information it has available.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetLatestNetworkInferencesOutlierResistant",
					Use:       "latest-network-inferences-outlier-resistant [topic_id]",
					Short:     "Get the latest Network inferences and weights for a topic. Will return whatever information it has available removing outliers.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetNetworkLossBundleAtBlock",
					Use:       "network-loss-bundle-at-block [topic_id] [block]",
					Short:     "Get the network loss bundle for a topic at given block height",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "IsWhitelistAdmin",
					Use:       "is-whitelist-admin [address]",
					Short:     "Check if an address is a whitelist admin. True if so, else false",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "GetStakeRemovalsUpUntilBlock",
					Use:       "stake-removals-up-until-block [block_height]",
					Short:     "Get all pending stake removal requests going to happen at a given block height",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetDelegateStakeRemovalsUpUntilBlock",
					Use:       "delegate-stake-removals-up-until-block [block_height]",
					Short:     "Get all pending delegate stake removal requests going to happen at a given block height",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetStakeRemovalInfo",
					Use:       "stake-removal-info [address] [topic_id]",
					Short:     "Get a pending stake removal for a reputer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "reputer"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetDelegateStakeRemovalInfo",
					Use:       "delegate-stake-removal-info [delegator] [reputer] [topic_id]",
					Short:     "Get a pending delegate stake removal for a delegator in a topic upon a reputer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "delegator"},
						{ProtoField: "reputer"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetTopicLastWorkerCommitInfo",
					Use:       "topic-last-worker-commit [topic_id]",
					Short:     "Get topic last commit by worker",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetTopicLastReputerCommitInfo",
					Use:       "topic-last-reputer-commit [topic_id]",
					Short:     "Get topic last commit by reputer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetActiveTopicsAtBlock",
					Use:       "active-topics-at-block [block_height]",
					Short:     "Get active topics at block",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "block_height"},
					},
				},
				{
					RpcMethod: "GetNextChurningBlockByTopicId",
					Use:       "next-churning-block [topic_id]",
					Short:     "Get next possible churning block by topic id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetActiveForecastersForTopic",
					Use:       "active-forecasters [topic_id]",
					Short:     "Get active forecasters for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetActiveInferersForTopic",
					Use:       "active-inferers [topic_id]",
					Short:     "Get active inferers for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetActiveReputersForTopic",
					Use:       "active-reputers [topic_id]",
					Short:     "Get active reputers for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				// add inclusion methods
				{
					RpcMethod: "GetCountInfererInclusionsInTopic",
					Use:       "count-inferer-inclusions-in-topic [topic_id] [inferer]",
					Short:     "Get count of inferer inclusions in topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "inferer"},
					},
				},
				{
					RpcMethod: "GetCountForecasterInclusionsInTopic",
					Use:       "count-forecaster-inclusions-in-topic [topic_id] [forecaster]",
					Short:     "Get count of forecaster inclusions in topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "forecaster"},
					},
				},
				{
					RpcMethod: "IsWhitelistedGlobalWorker",
					Use:       "is-whitelisted-global-worker [address]",
					Short:     "Check if address is whitelisted as global worker",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsWhitelistedGlobalReputer",
					Use:       "is-whitelisted-global-reputer [address]",
					Short:     "Check if address is whitelisted as global reputer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsWhitelistedGlobalAdmin",
					Use:       "is-whitelisted-global-admin [address]",
					Short:     "Check if address is whitelisted as global admin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanUpdateAllGlobalWhitelists",
					Use:       "can-update-all-global-whitelists [address]",
					Short:     "Check if address can update all global whitelists",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanUpdateGlobalWorkerWhitelist",
					Use:       "can-update-global-worker-whitelist [address]",
					Short:     "Check if address can update global worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanUpdateGlobalReputerWhitelist",
					Use:       "can-update-global-reputer-whitelist [address]",
					Short:     "Check if address can update global reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsTopicWorkerWhitelistEnabled",
					Use:       "is-topic-worker-whitelist-enabled [topic_id]",
					Short:     "Check if topic-level worker whitelist is enabled for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "IsTopicReputerWhitelistEnabled",
					Use:       "is-topic-reputer-whitelist-enabled [topic_id]",
					Short:     "Check if topic-level reputer whitelist is enabled for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "IsWhitelistedTopicCreator",
					Use:       "is-whitelisted-topic-creator [address]",
					Short:     "Check if address is whitelisted as topic creator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsWhitelistedGlobalActor",
					Use:       "is-whitelist-global-actor [address]",
					Short:     "Check if address is a member of the global whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsWhitelistedTopicWorker",
					Use:       "is-whitelisted-topic-worker [topic_id] [address]",
					Short:     "Check if address is whitelisted as worker for topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "IsWhitelistedTopicReputer",
					Use:       "is-whitelisted-topic-reputer [topic_id] [address]",
					Short:     "Check if address is whitelisted as reputer for topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanUpdateAllGlobalWhitelists",
					Use:       "can-update-global-whitelists [address]",
					Short:     "Check if address can update global whitelists",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanUpdateParams",
					Use:       "can-update-params [address]",
					Short:     "Check if address can update params",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanUpdateTopicWhitelist",
					Use:       "can-update-topic-whitelist [topic_id] [address]",
					Short:     "Check if address can update topic whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanCreateTopic",
					Use:       "can-create-topic [address]",
					Short:     "Check if address can create topics",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanSubmitWorkerPayload",
					Use:       "can-submit-worker-payload [topic_id] [address]",
					Short:     "Check if address can submit worker payload for topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "CanSubmitReputerPayload",
					Use:       "can-submit-reputer-payload [topic_id] [address]",
					Short:     "Check if address can submit reputer payload for topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "GetTopicInitialInfererEmaScore",
					Use:       "initial-inferer-ema-score [topic_id]",
					Short:     "Get the initial EMA score for inferers in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetTopicInitialForecasterEmaScore",
					Use:       "initial-forecaster-ema-score [topic_id]",
					Short:     "Get the initial EMA score for forecasters in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "GetTopicInitialReputerEmaScore",
					Use:       "initial-reputer-ema-score [topic_id]",
					Short:     "Get the initial EMA score for reputers in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "topic_id"},
					},
				},
			},
			SubCommands:          nil,
			EnhanceCustomCommand: false,
			Short:                "Emissions module query commands",
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: statev7.MsgService_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Use:       "update-params [sender] [params]",
					Short:     "Update params of the network",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "params"},
					},
				},
				{
					RpcMethod: "CreateNewTopic",
					Use:       "create-topic [creator] [metadata] [loss_method] [epoch_length] [ground_truth_lag] [worker_submission_window] [p_norm] [alpha_regret] [allow_negative] [epsilon] [merit_sortition_alpha] [active_inferer_quantile] [active_forecaster_quantile] [active_reputer_quantile] [enable_worker_whitelist] [enable_reputer_whitelist]",
					Short:     "Add a new topic to the network",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "creator"},
						{ProtoField: "metadata"},
						{ProtoField: "loss_method"},
						{ProtoField: "epoch_length"},
						{ProtoField: "ground_truth_lag"},
						{ProtoField: "worker_submission_window"},
						{ProtoField: "p_norm"},
						{ProtoField: "alpha_regret"},
						{ProtoField: "allow_negative"},
						{ProtoField: "epsilon"},
						{ProtoField: "merit_sortition_alpha"},
						{ProtoField: "active_inferer_quantile"},
						{ProtoField: "active_forecaster_quantile"},
						{ProtoField: "active_reputer_quantile"},
						{ProtoField: "enable_worker_whitelist"},
						{ProtoField: "enable_reputer_whitelist"},
					},
				},
				{
					RpcMethod: "Register",
					Use:       "register [sender] [topic_ids] [owner] [is_reputer]",
					Short:     "Register a new reputer or worker for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "owner"},
						{ProtoField: "is_reputer"},
					},
				},
				{
					RpcMethod: "RemoveRegistration",
					Use:       "remove-registration [sender] [owner] [is_reputer]",
					Short:     "Remove a reputer or worker from a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "is_reputer"},
					},
				},
				{
					RpcMethod: "AddStake",
					Use:       "add-stake [sender] [topic_id] [amount]",
					Short:     "Add stake [amount] to ones self sender [reputer or worker] for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "RemoveStake",
					Use:       "remove-stake [sender] [topic_id] [amount]",
					Short:     "modify sender's [reputer] stake position by removing [amount] stake from a topic [topic_id]",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "CancelRemoveStake",
					Use:       "cancel-remove-stake [sender] [topic_id]",
					Short:     "Cancel the removal of stake for a reputer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "DelegateStake",
					Use:       "delegate-stake [sender] [topic_id] [reputer] [amount]",
					Short:     "Delegate stake [amount] to a reputer for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "RemoveDelegateStake",
					Use:       "remove-delegate-stake [sender] [topic_id] [reputer] [amount]",
					Short:     "Modify sender's [reputer] delegate stake position by removing [amount] stake from a topic [topic_id] from a reputer [reputer]",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "CancelRemoveDelegateStake",
					Use:       "cancel-remove-delegate-stake [sender] [topic_id] [reputer]",
					Short:     "Cancel the removal of delegated stake for a delegator staking on a reputer in a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "RewardDelegateStake",
					Use:       "reward-delegate-stake [sender] [topic_id] [reputer]",
					Short:     "Get Reward for Delegator [sender] for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "reputer"},
					},
				},
				{
					RpcMethod: "FundTopic",
					Use:       "fund-topic [sender] [topic_id] [amount] [extra_data]",
					Short:     "send funds to a topic to pay for inferences",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "AddToWhitelistAdmin",
					Use:       "add-to-whitelist-admin [sender] [address]",
					Short:     "add an admin address to the whitelist used for admin functions on-chain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "RemoveFromWhitelistAdmin",
					Use:       "remove-from-whitelist-admin [sender] [address]",
					Short:     "remove a admin address from the whitelist used for admin functions on-chain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "InsertWorkerPayload",
					Use:       "insert-worker-payload [sender] [worker_data]",
					Short:     "Insert worker payload",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "worker_data_bundle"},
					},
				},
				{
					RpcMethod: "InsertReputerPayload",
					Use:       "insert-reputer-payload [sender] [reputer_data]",
					Short:     "Insert reputer payload",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "reputer_value_bundle"},
					},
				},
				{
					RpcMethod: "AddToGlobalWhitelist",
					Use:       "add-to-global-whitelist [sender] [address]",
					Short:     "Add an address to the global whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "AddToGlobalWorkerWhitelist",
					Use:       "add-to-global-worker-whitelist [sender] [address]",
					Short:     "Add an address to the global worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "RemoveFromGlobalWorkerWhitelist",
					Use:       "remove-from-global-worker-whitelist [sender] [address]",
					Short:     "Remove an address from the global worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "AddToGlobalReputerWhitelist",
					Use:       "add-to-global-reputer-whitelist [sender] [address]",
					Short:     "Add an address to the global reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "RemoveFromGlobalReputerWhitelist",
					Use:       "remove-from-global-reputer-whitelist [sender] [address]",
					Short:     "Remove an address from the global reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "AddToGlobalAdminWhitelist",
					Use:       "add-to-global-admin-whitelist [sender] [address]",
					Short:     "Add an address to the global admin whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "RemoveFromGlobalAdminWhitelist",
					Use:       "remove-from-global-admin-whitelist [sender] [address]",
					Short:     "Remove an address from the global admin whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "BulkAddToGlobalWorkerWhitelist",
					Use:       "bulk-add-to-global-worker-whitelist [sender] [addresses]",
					Short:     "Bulk add addresses to the global worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "BulkRemoveFromGlobalWorkerWhitelist",
					Use:       "bulk-remove-from-global-worker-whitelist [sender] [addresses]",
					Short:     "Bulk remove addresses from the global worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "BulkAddToGlobalReputerWhitelist",
					Use:       "bulk-add-to-global-reputer-whitelist [sender] [addresses]",
					Short:     "Bulk add addresses to the global reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "BulkRemoveFromGlobalReputerWhitelist",
					Use:       "bulk-remove-from-global-reputer-whitelist [sender] [addresses]",
					Short:     "Bulk remove addresses from the global reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "BulkAddToTopicWorkerWhitelist",
					Use:       "bulk-add-to-topic-worker-whitelist [sender] [topic_id] [addresses]",
					Short:     "Bulk add addresses to a topic's worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "BulkRemoveFromTopicWorkerWhitelist",
					Use:       "bulk-remove-from-topic-worker-whitelist [sender] [topic_id] [addresses]",
					Short:     "Bulk remove addresses from a topic's worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "BulkAddToTopicReputerWhitelist",
					Use:       "bulk-add-to-topic-reputer-whitelist [sender] [topic_id] [addresses]",
					Short:     "Bulk add addresses to a topic's reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "BulkRemoveFromTopicReputerWhitelist",
					Use:       "bulk-remove-from-topic-reputer-whitelist [sender] [topic_id] [addresses]",
					Short:     "Bulk remove addresses from a topic's reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
						{ProtoField: "addresses"},
					},
				},
				{
					RpcMethod: "RemoveFromGlobalWhitelist",
					Use:       "remove-from-global-whitelist [sender] [address]",
					Short:     "Remove an address from the global whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "EnableTopicWorkerWhitelist",
					Use:       "enable-topic-worker-whitelist [sender] [topic_id]",
					Short:     "Enable worker whitelist for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "DisableTopicWorkerWhitelist",
					Use:       "disable-topic-worker-whitelist [sender] [topic_id]",
					Short:     "Disable worker whitelist for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "EnableTopicReputerWhitelist",
					Use:       "enable-topic-reputer-whitelist [sender] [topic_id]",
					Short:     "Enable reputer whitelist for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "DisableTopicReputerWhitelist",
					Use:       "disable-topic-reputer-whitelist [sender] [topic_id]",
					Short:     "Disable reputer whitelist for a topic",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "AddToTopicCreatorWhitelist",
					Use:       "add-to-topic-creator-whitelist [sender] [address]",
					Short:     "Add an address to the topic creator whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "RemoveFromTopicCreatorWhitelist",
					Use:       "remove-from-topic-creator-whitelist [sender] [address]",
					Short:     "Remove an address from the topic creator whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "AddToTopicWorkerWhitelist",
					Use:       "add-to-topic-worker-whitelist [sender] [address] [topic_id]",
					Short:     "Add an address to a topic's worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "RemoveFromTopicWorkerWhitelist",
					Use:       "remove-from-topic-worker-whitelist [sender] [address] [topic_id]",
					Short:     "Remove an address from a topic's worker whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "AddToTopicReputerWhitelist",
					Use:       "add-to-topic-reputer-whitelist [sender] [address] [topic_id]",
					Short:     "Add an address to a topic's reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
						{ProtoField: "topic_id"},
					},
				},
				{
					RpcMethod: "RemoveFromTopicReputerWhitelist",
					Use:       "remove-from-topic-reputer-whitelist [sender] [address] [topic_id]",
					Short:     "Remove an address from a topic's reputer whitelist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
						{ProtoField: "address"},
						{ProtoField: "topic_id"},
					},
				},
			},
			SubCommands:          nil,
			EnhanceCustomCommand: false,
			Short:                "Emissions module transaction commands",
		},
	}
}
