package stress_test

import (
	"os"
	"runtime"
	"testing"

	testCommon "github.com/allora-network/allora-chain/test/common"
)

func TestStressTestSuite(t *testing.T) {
	if _, isIntegration := os.LookupEnv("STRESS_TEST"); isIntegration == false {
		t.Skip("Skipping Stress Test unless explicitly enabled")
	}

	numCPUs := runtime.NumCPU()
	gomaxprocs := runtime.GOMAXPROCS(0)
	t.Logf("Number of logical CPUs: %d, GOMAXPROCS %d \n", numCPUs, gomaxprocs)

	t.Log(">>> Setting up connection to local node <<<")

	seed := testCommon.LookupEnvIntWithDefault(t, "SEED", 0)
	rpcMode := testCommon.LookupRpcModeWithDefault(t, "RPC_MODE", testCommon.SingleRpc)
	rpcEndpoints := testCommon.LookupEnvStringArrayWithDefault(t, "RPC_URLS", []string{"http://localhost:26657"})

	testConfig := testCommon.NewTestConfig(
		t,
		rpcMode,
		rpcEndpoints,
		"../localnet/genesis",
		seed,
	)

	// Read env vars with defaults
	reputersPerIteration := testCommon.LookupEnvIntWithDefault(t, "REPUTERS_PER_ITERATION", 1)
	maxReputersPerTopic := testCommon.LookupEnvIntWithDefault(t, "MAX_REPUTERS_PER_TOPIC", 20)
	workersPerIteration := testCommon.LookupEnvIntWithDefault(t, "WORKERS_PER_ITERATION", 1)
	maxWorkersPerTopic := testCommon.LookupEnvIntWithDefault(t, "MAX_WORKERS_PER_TOPIC", 20)
	topicsPerIteration := testCommon.LookupEnvIntWithDefault(t, "TOPICS_PER_ITERATION", 1)
	topicsMax := testCommon.LookupEnvIntWithDefault(t, "TOPICS_MAX", 100)
	maxIterations := testCommon.LookupEnvIntWithDefault(t, "MAX_ITERATIONS", 1000)
	epochLength := testCommon.LookupEnvIntWithDefault(t, "EPOCH_LENGTH", 12)
	doFinalReport := testCommon.LookupEnvBoolWithDefault(t, "FINAL_REPORT", false)

	t.Log("Reputers per iteration: ", reputersPerIteration)
	t.Log("Max Reputers per topic: ", maxReputersPerTopic)
	t.Log("Workers per iteration: ", workersPerIteration)
	t.Log("Max Workers per topic: ", maxWorkersPerTopic)
	t.Log("Topics per iteration of topics: ", topicsPerIteration)
	t.Log("Topics global max: ", topicsMax)
	t.Log("Max worker+reputer iterations: ", maxIterations)
	t.Log("Epoch Length: ", epochLength)
	t.Log("Use mutex to prepare final report: ", doFinalReport)

	t.Log(">>> Test Making Inference <<<")
	workerReputerCoordinationLoop(
		testConfig,
		reputersPerIteration,
		maxReputersPerTopic,
		workersPerIteration,
		maxWorkersPerTopic,
		topicsPerIteration,
		topicsMax,
		maxIterations,
		epochLength,
		doFinalReport,
	)
}
