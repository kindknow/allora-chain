package integration_test

import (
	"context"
	"fmt"

	cosmosMath "cosmossdk.io/math"
	testCommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

const STAKING_TOPIC_ID = uint64(1)

func addTopicReputer(m testCommon.TestConfig, address string, topicId uint64) {
	ctx := context.Background()
	addTopicReputerRequest := &emissionstypes.AddToTopicReputerWhitelistRequest{
		Sender:  m.AliceAddr,
		Address: address,
		TopicId: topicId,
	}
	txResp, err := m.Client.BroadcastTx(ctx, m.AliceAcc, addTopicReputerRequest)
	require.NoError(m.T, err)
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)
	addTopicReputerResponse := &emissionstypes.AddToTopicReputerWhitelistResponse{}
	err = txResp.Decode(addTopicReputerResponse)
	require.NoError(m.T, err)
}

// register alice as a reputer in topic 1, then check success
func StakeAliceAsReputerTopic1(m testCommon.TestConfig) {
	ctx := context.Background()
	// Record Alice stake before adding more
	aliceStakedBefore, err := m.Client.QueryEmissions().GetReputerStakeInTopic(
		ctx,
		&emissionstypes.GetReputerStakeInTopicRequest{
			TopicId: STAKING_TOPIC_ID,
			Address: m.AliceAddr,
		},
	)
	require.NoError(m.T, err)

	const stakeToAdd = 1000000

	// Allow Alice to stake as a reputer
	addTopicReputer(m, m.AliceAddr, STAKING_TOPIC_ID)

	// Have Alice stake more
	addStake := &emissionstypes.AddStakeRequest{
		Sender:  m.AliceAddr,
		TopicId: STAKING_TOPIC_ID,
		Amount:  cosmosMath.NewInt(stakeToAdd),
	}
	txResp, err := m.Client.BroadcastTx(ctx, m.AliceAcc, addStake)
	require.NoError(m.T, err)
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	// Check Alice has stake on the topic
	aliceStakedAfter, err := m.Client.QueryEmissions().GetReputerStakeInTopic(
		ctx,
		&emissionstypes.GetReputerStakeInTopicRequest{
			TopicId: STAKING_TOPIC_ID,
			Address: m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.Equal(m.T, fmt.Sprint(stakeToAdd), aliceStakedAfter.Amount.Sub(aliceStakedBefore.Amount).String())
}

// integration tests the ability of bob to stake on alice as a reputer
func StakeBobOnAliceAsReputerTopic1(m testCommon.TestConfig) {
	ctx := context.Background()
	// Record Bob stake before adding more
	bobStakedBefore, err := m.Client.QueryEmissions().GetStakeFromDelegatorInTopicInReputer(
		ctx,
		&emissionstypes.GetStakeFromDelegatorInTopicInReputerRequest{
			TopicId:          STAKING_TOPIC_ID,
			DelegatorAddress: m.BobAddr,
			ReputerAddress:   m.AliceAddr,
		},
	)
	require.NoError(m.T, err)

	const stakeToAdd = 200000

	// Have bob stake
	addDelegateStake := &emissionstypes.DelegateStakeRequest{
		Sender:  m.BobAddr,
		Reputer: m.AliceAddr,
		TopicId: STAKING_TOPIC_ID,
		Amount:  cosmosMath.NewInt(stakeToAdd),
	}
	txResp, err := m.Client.BroadcastTx(ctx, m.BobAcc, addDelegateStake)
	require.NoError(m.T, err)
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	// Check Alice has stake on the topic
	bobStakedAfter, err := m.Client.QueryEmissions().GetStakeFromDelegatorInTopicInReputer(
		ctx,
		&emissionstypes.GetStakeFromDelegatorInTopicInReputerRequest{
			TopicId:          STAKING_TOPIC_ID,
			DelegatorAddress: m.BobAddr,
			ReputerAddress:   m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.Equal(m.T, fmt.Sprint(stakeToAdd), bobStakedAfter.Amount.Sub(bobStakedBefore.Amount).String())
}

// Register two actors and check their registrations went through
func StakingChecks(m testCommon.TestConfig) {
	ctx := context.Background()
	m.T.Log("--- Staking Alice as Reputer ---")
	StakeAliceAsReputerTopic1(m)

	res, _ := m.Client.QueryEmissions().GetTopic(ctx, &emissionstypes.GetTopicRequest{
		TopicId: STAKING_TOPIC_ID,
	})
	// Topic is not expected to be funded yet => expect 0 weight => topic not active!
	// But we still have this conditional just in case there are > 0 funds
	if res.EffectiveRevenue != "0" {
		m.T.Log("--- Check reactivating Topic 1 ---")
		CheckTopic1Activated(m)
	}

	m.T.Log("--- Staking Bob on Alice as Reputer ---")
	StakeBobOnAliceAsReputerTopic1(m)
}

// Unstake Alice as a reputer in topic 1, then check success
func UnstakeAliceAsReputerTopic1(m testCommon.TestConfig) {
	ctx := context.Background()
	aliceStakeBefore, err := m.Client.QueryEmissions().GetStakeFromReputerInTopicInSelf(
		ctx,
		&emissionstypes.GetStakeFromReputerInTopicInSelfRequest{
			TopicId:        STAKING_TOPIC_ID,
			ReputerAddress: m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.True(
		m.T,
		aliceStakeBefore.Amount.GT(cosmosMath.ZeroInt()),
		"Alice should have stake in topic 1",
	)

	// Have Alice unstake
	unstake := &emissionstypes.RemoveStakeRequest{
		Sender:  m.AliceAddr,
		TopicId: STAKING_TOPIC_ID,
		Amount:  aliceStakeBefore.Amount,
	}

	txResp, err := m.Client.BroadcastTx(ctx, m.AliceAcc, unstake)
	require.NoError(m.T, err)
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	// check the unstake removal is queued
	stakeRemoval, err := m.Client.QueryEmissions().GetStakeRemovalInfo(
		ctx,
		&emissionstypes.GetStakeRemovalInfoRequest{
			TopicId: STAKING_TOPIC_ID,
			Reputer: m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.NotNil(m.T, stakeRemoval)
	require.NotZero(m.T, stakeRemoval.Removal.BlockRemovalCompleted)
	m.T.Log("--- Unstake removal is queued, waiting for block ", stakeRemoval.Removal.BlockRemovalCompleted, " ---")
	err = m.Client.WaitForBlockHeight(ctx, stakeRemoval.Removal.BlockRemovalCompleted+1)
	require.NoError(m.T, err)
	blockHeight, err := m.Client.BlockHeight(ctx)
	require.NoError(m.T, err)
	require.Greater(m.T, blockHeight, stakeRemoval.Removal.BlockRemovalCompleted)

	// Check Alice has zero stake left
	aliceStakedAfter, err := m.Client.QueryEmissions().GetStakeFromReputerInTopicInSelf(
		ctx,
		&emissionstypes.GetStakeFromReputerInTopicInSelfRequest{
			TopicId:        STAKING_TOPIC_ID,
			ReputerAddress: m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.True(
		m.T,
		aliceStakedAfter.Amount.Equal(cosmosMath.ZeroInt()),
		"Alice should have zero stake in topic 1 after unstake",
		stakeRemoval.Removal,
		aliceStakeBefore.Amount.String(),
		aliceStakedAfter.Amount.String(),
	)
}

// Unstake Bob as a delegator delegated to Alice in topic 1, then check success
func UnstakeBobAsDelegatorOnAliceTopic1(m testCommon.TestConfig) {
	ctx := context.Background()
	bobStake, err := m.Client.QueryEmissions().GetStakeFromDelegatorInTopicInReputer(
		ctx,
		&emissionstypes.GetStakeFromDelegatorInTopicInReputerRequest{
			TopicId:          STAKING_TOPIC_ID,
			DelegatorAddress: m.BobAddr,
			ReputerAddress:   m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.True(
		m.T,
		bobStake.Amount.GT(cosmosMath.ZeroInt()),
		"Bob should have stake on Alice in topic 1",
	)

	// Have Bob unstake
	unstake := &emissionstypes.RemoveDelegateStakeRequest{
		Sender:  m.BobAddr,
		Reputer: m.AliceAddr,
		TopicId: STAKING_TOPIC_ID,
		Amount:  bobStake.Amount,
	}

	txResp, err := m.Client.BroadcastTx(ctx, m.BobAcc, unstake)
	require.NoError(m.T, err)
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	// check the unstake removal is queued
	stakeRemoval, err := m.Client.QueryEmissions().GetDelegateStakeRemovalInfo(
		ctx,
		&emissionstypes.GetDelegateStakeRemovalInfoRequest{
			TopicId:   STAKING_TOPIC_ID,
			Delegator: m.BobAddr,
			Reputer:   m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.NotNil(m.T, stakeRemoval)
	require.NotZero(m.T, stakeRemoval.Removal.BlockRemovalCompleted)
	m.T.Log("--- Unstake removal is queued, waiting for block ", stakeRemoval.Removal.BlockRemovalCompleted, " ---")
	err = m.Client.WaitForBlockHeight(ctx, stakeRemoval.Removal.BlockRemovalCompleted+1)
	require.NoError(m.T, err)

	// Check Bob has zero stake left
	bobStakedAfter, err := m.Client.QueryEmissions().GetStakeFromDelegatorInTopicInReputer(
		ctx,
		&emissionstypes.GetStakeFromDelegatorInTopicInReputerRequest{
			TopicId:          STAKING_TOPIC_ID,
			DelegatorAddress: m.BobAddr,
			ReputerAddress:   m.AliceAddr,
		},
	)
	require.NoError(m.T, err)
	require.True(
		m.T,
		bobStakedAfter.Amount.Equal(cosmosMath.ZeroInt()),
		"Bob should have zero stake in topic 1 after unstake",
	)
}

// run checks for unstaking
func UnstakingChecks(m testCommon.TestConfig) {
	m.T.Log("--- Bob Unstaking as Delegator Upon Alice ---")
	UnstakeBobAsDelegatorOnAliceTopic1(m)
	m.T.Log("--- Unstaking Alice as Reputer ---")
	UnstakeAliceAsReputerTopic1(m)
}
