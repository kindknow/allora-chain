package app

import (
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/allora-network/allora-chain/app/upgrades"
	"github.com/allora-network/allora-chain/app/upgrades/v0_3_0"
	"github.com/allora-network/allora-chain/app/upgrades/v0_4_0"
	"github.com/allora-network/allora-chain/app/upgrades/v0_5_0"
	"github.com/allora-network/allora-chain/app/upgrades/v0_6_0"
	"github.com/allora-network/allora-chain/app/upgrades/v0_7_0"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"
)

var upgradeHandlers = []upgrades.Upgrade{
	v0_3_0.Upgrade,
	v0_4_0.Upgrade,
	v0_5_0.Upgrade,
	v0_6_0.Upgrade,
	v0_7_0.Upgrade,
	// Add more upgrade handlers here
	// ...
}

func (app *AlloraApp) setupUpgradeHandlers() {
	for _, handler := range upgradeHandlers {
		app.UpgradeKeeper.SetUpgradeHandler(handler.UpgradeName,
			handler.CreateUpgradeHandler(app.ModuleManager, app.Configurator()))

		upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
		if err != nil {
			panic(err)
		}

		if upgradeInfo.Name == v0_7_0.Upgrade.UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
			storeUpgrades := storetypes.StoreUpgrades{
				Added: []string{feemarkettypes.ModuleName},
			}

			// configure store loader that checks if version == upgradeHeight and applies store upgrades
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
		}
	}
}
