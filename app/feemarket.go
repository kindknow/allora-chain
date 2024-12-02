package app

import (
	"fmt"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	feemarket "github.com/skip-mev/feemarket/x/feemarket"
	feemarketkeeper "github.com/skip-mev/feemarket/x/feemarket/keeper"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"
)

// registerFeeMarketModule registers the feemarket module.
func (app *AlloraApp) registerFeeMarketModule() {
	if err := app.RegisterStores(
		storetypes.NewKVStoreKey(feemarkettypes.StoreKey),
	); err != nil {
		panic(err)
	}

	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		app.appCodec,
		app.GetKey(feemarkettypes.StoreKey),
		app.AccountKeeper,
		&DefaultFeemarketDenomResolver{},
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	if err := app.RegisterModules(
		feemarket.NewAppModule(app.appCodec, *app.FeeMarketKeeper),
	); err != nil {
		panic(err)
	}
}

type DefaultFeemarketDenomResolver struct{}

func (r *DefaultFeemarketDenomResolver) ConvertToDenom(_ sdk.Context, coin sdk.DecCoin, denom string) (sdk.DecCoin, error) {
	if coin.Denom == denom {
		return coin, nil
	}

	return sdk.DecCoin{}, fmt.Errorf("error resolving denom: the only denom supported is %s", coin.Denom)
}

func (r *DefaultFeemarketDenomResolver) ExtraDenoms(_ sdk.Context) ([]string, error) {
	return []string{}, nil
}
