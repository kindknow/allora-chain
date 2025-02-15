package stress_test

import (
	"context"

	cosmosMath "cosmossdk.io/math"
	alloraMath "github.com/allora-network/allora-chain/math"
	testCommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

// Broadcast the tx to create a new topic
func createTopic(
	m testCommon.TestConfig,
	epochLength int64,
	creator NameAccountAndAddress,
) (topicId uint64) {
	ctx := context.Background()
	createTopicRequest := &emissionstypes.CreateNewTopicRequest{
		Creator:                  creator.aa.addr,
		Metadata:                 "ETH 24h Prediction",
		LossMethod:               "mse",
		EpochLength:              epochLength,
		GroundTruthLag:           epochLength,
		WorkerSubmissionWindow:   epochLength,
		PNorm:                    alloraMath.NewDecFromInt64(3),
		AlphaRegret:              alloraMath.NewDecFromInt64(1),
		AllowNegative:            true,
		Epsilon:                  alloraMath.MustNewDecFromString("0.01"),
		MeritSortitionAlpha:      alloraMath.MustNewDecFromString("0.1"),
		ActiveInfererQuantile:    alloraMath.MustNewDecFromString("0.2"),
		ActiveForecasterQuantile: alloraMath.MustNewDecFromString("0.2"),
		ActiveReputerQuantile:    alloraMath.MustNewDecFromString("0.2"),
		EnableWorkerWhitelist:    true,
		EnableReputerWhitelist:   true,
	}

	txResp, err := m.Client.BroadcastTx(ctx, creator.aa.acc, createTopicRequest)
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	createTopicResponse := &emissionstypes.CreateNewTopicResponse{} //nolint:exhaustruct // the fields are populated by decode
	err = txResp.Decode(createTopicResponse)
	require.NoError(m.T, err)

	incrementCountTopics()

	m.T.Log(topicLog(createTopicResponse.TopicId, "creator", creator.name, "created topic"))

	return createTopicResponse.TopicId
}

// broadcast a tx to fund a topic
func fundTopic(
	m testCommon.TestConfig,
	topicId uint64,
	sender NameAccountAndAddress,
	amount int64,
) error {
	ctx := context.Background()
	m.T.Log(topicLog(topicId, "funded topic with ", amount, "from", sender.name))
	txResp, err := m.Client.BroadcastTx(
		ctx,
		sender.aa.acc,
		&emissionstypes.FundTopicRequest{
			Sender:  sender.aa.addr,
			TopicId: topicId,
			Amount:  cosmosMath.NewInt(amount),
		},
	)
	if err != nil {
		return err
	}
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	if err != nil {
		return err
	}
	resp := &emissionstypes.FundTopicResponse{}
	err = txResp.Decode(resp)
	if err != nil {
		return err
	}
	return nil
}
