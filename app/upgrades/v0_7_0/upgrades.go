package v0_7_0 //nolint:revive // var-naming: don't use an underscore in package name

import (
	"context"

	cosmosmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/allora-network/allora-chain/app/keepers"
	"github.com/allora-network/allora-chain/app/params"
	"github.com/allora-network/allora-chain/app/upgrades"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"
)

const (
	UpgradeName = "v0.7.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			feegrant.StoreKey, feemarkettypes.StoreKey,
		},
		Renamed: nil,
		Deleted: nil,
	},
}

func CreateUpgradeHandler(
	moduleManager *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {

		vm, err := moduleManager.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		if err := ConfigureFeeMarketModule(ctx, keepers); err != nil {
			return vm, err
		}

		return vm, nil
	}
}

func ConfigureFeeMarketModule(ctx context.Context, keepers *keepers.AppKeepers) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	feeMarketParams, err := keepers.FeeMarketKeeper.GetParams(sdkCtx)
	if err != nil {
		return err
	}

	feeMarketParams.Enabled = true
	feeMarketParams.FeeDenom = params.BaseCoinUnit
	feeMarketParams.DistributeFees = true
	feeMarketParams.MinBaseGasPrice = cosmosmath.LegacyMustNewDecFromStr("10")
	if err := keepers.FeeMarketKeeper.SetParams(sdkCtx, feeMarketParams); err != nil {
		return err
	}

	state, err := keepers.FeeMarketKeeper.GetState(sdkCtx)
	if err != nil {
		return err
	}

	state.BaseGasPrice = cosmosmath.LegacyMustNewDecFromStr("10")

	return keepers.FeeMarketKeeper.SetState(sdkCtx, state)
}
