package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/allora-network/allora-chain/app/params"
	alloraMath "github.com/allora-network/allora-chain/math"
	"github.com/allora-network/allora-chain/x/mint/types"

	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeService     storetypes.KVStoreService
	accountKeeper    types.AccountKeeper
	stakingKeeper    types.StakingKeeper
	bankKeeper       types.BankKeeper
	emissionsKeeper  types.EmissionsKeeper
	feeCollectorName string

	Schema                                   collections.Schema
	Params                                   collections.Item[types.Params]
	PreviousRewardEmissionPerUnitStakedToken collections.Item[math.LegacyDec]
	PreviousBlockEmission                    collections.Item[math.Int]
	EcosystemTokensMinted                    collections.Item[math.Int]
	MonthsUnlocked                           collections.Item[math.Int]
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	sk types.StakingKeeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	ek types.EmissionsKeeper,
	feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("the x/%s module account has not been set", types.ModuleName))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		Schema:                                   collections.Schema{},
		cdc:                                      cdc,
		storeService:                             storeService,
		stakingKeeper:                            sk,
		accountKeeper:                            ak,
		bankKeeper:                               bk,
		emissionsKeeper:                          ek,
		feeCollectorName:                         feeCollectorName,
		Params:                                   collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		PreviousRewardEmissionPerUnitStakedToken: collections.NewItem(sb, types.PreviousRewardEmissionPerUnitStakedTokenKey, "previousrewardsemissionsperunitstakedtoken", alloraMath.LegacyDecValue),
		PreviousBlockEmission:                    collections.NewItem(sb, types.PreviousBlockEmissionKey, "previousblockemission", sdk.IntValue),
		EcosystemTokensMinted:                    collections.NewItem(sb, types.EcosystemTokensMintedKey, "ecosystemtokensminted", sdk.IntValue),
		MonthsUnlocked:                           collections.NewItem(sb, types.MonthsUnlockedKey, "monthsunlocked", sdk.IntValue),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
	return k
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", "x/"+types.ModuleName)
}

// getter for the storage service
func (k Keeper) GetStorageService() storetypes.KVStoreService {
	return k.storeService
}

// getter for the binary codec
func (k Keeper) GetBinaryCodec() codec.BinaryCodec {
	return k.cdc
}

// This function increases the ledger that tracks the total tokens minted by the ecosystem treasury
// over the life of the blockchain.
func (k Keeper) AddEcosystemTokensMinted(ctx context.Context, minted math.Int) error {
	curr, err := k.EcosystemTokensMinted.Get(ctx)
	if err != nil {
		return err
	}
	newTotal := curr.Add(minted)
	return k.EcosystemTokensMinted.Set(ctx, newTotal)
}

// Setter for the number of months unlocked
// this function coerces values to be between 0 and 36
func (k Keeper) SetMonthsAlreadyUnlocked(ctx context.Context, months math.Int) error {
	if months.IsNegative() {
		months = math.ZeroInt()
	}
	if months.GT(math.NewInt(36)) {
		months = math.NewInt(36)
	}
	return k.MonthsUnlocked.Set(ctx, months)
}

/// STAKING KEEPER RELATED FUNCTIONS

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) CosmosValidatorStakedSupply(ctx context.Context) (math.Int, error) {
	return k.stakingKeeper.TotalBondedTokens(ctx)
}

/// BANK KEEPER RELATED FUNCTIONS

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx context.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// MoveCoinsFromMintToEcosystem moves freshly minted tokens from the mint module
// which has permissions to create new tokens, to the ecosystem account which
// only has permissions to hold tokens.
func (k Keeper) MoveCoinsFromMintToEcosystem(ctx context.Context, mintedCoins sdk.Coins) error {
	if mintedCoins.Empty() {
		return nil
	}
	return k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.ModuleName,
		types.EcosystemModuleName,
		mintedCoins,
	)
}

// PayValidatorsFromEcosystem sends funds from the ecosystem
// treasury account to the cosmos network validators rewards account (fee collector)
// PayValidatorsFromEcosystem to be used in BeginBlocker.
func (k Keeper) PayValidatorsFromEcosystem(ctx context.Context, rewards sdk.Coins) error {
	if rewards.Empty() {
		return nil
	}
	return k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.EcosystemModuleName,
		k.feeCollectorName,
		rewards,
	)
}

// PayAlloraRewardsFromEcosystem sends funds from the ecosystem
// treasury account to the allora reward payout account used in the emissions module
// PayAlloraRewardsFromEcosystem to be used in BeginBlocker.
func (k Keeper) PayAlloraRewardsFromEcosystem(ctx context.Context, rewards sdk.Coins) error {
	if rewards.Empty() {
		return nil
	}
	err := k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.EcosystemModuleName,
		emissionstypes.AlloraRewardsAccountName,
		rewards,
	)
	if err != nil {
		return err
	}

	return k.emissionsKeeper.SetRewardCurrentBlockEmission(ctx, rewards.AmountOf(params.BaseCoinUnit))
}

// GetTotalCurrTokenSupply implements an alias call to the underlying supply keeper's
// GetTotalCurrTokenSupply to be used in BeginBlocker.
func (k Keeper) GetTotalCurrTokenSupply(ctx context.Context) sdk.Coin {
	return k.bankKeeper.GetSupply(ctx, params.BaseCoinUnit)
}

// returns the quantity of tokens currently stored in the "ecosystem" module account
// this module account is paid by inference requests and is drained by this mint module
// when forwarding rewards to fee collector and allorarewards accounts
func (k Keeper) GetEcosystemBalance(ctx context.Context, mintDenom string) (math.Int, error) {
	ecosystemAddr := k.accountKeeper.GetModuleAddress(types.EcosystemModuleName)
	return k.bankKeeper.GetBalance(ctx, ecosystemAddr, mintDenom).Amount, nil
}

// Params getter
func (k Keeper) GetParams(ctx context.Context) (types.Params, error) {
	return k.Params.Get(ctx)
}

// What split of the rewards should be given to cosmos validators vs
// allora participants (reputers, forecaster workers, inferrer workers)
func (k Keeper) GetValidatorsVsAlloraPercentReward(ctx context.Context) (alloraMath.Dec, error) {
	emissionsParams, err := k.emissionsKeeper.GetParams(ctx)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	return emissionsParams.ValidatorsVsAlloraPercentReward, nil
}

// The last time we paid out rewards, what was the percentage of those rewards that went to staked reputers
// (as opposed to forecaster workers and inferrer workers)
func (k Keeper) GetPreviousPercentageRewardToStakedReputers(ctx context.Context) (math.LegacyDec, error) {
	stakedPercent, err := k.emissionsKeeper.GetPreviousPercentageRewardToStakedReputers(ctx)
	if err != nil {
		return math.LegacyDec{}, err
	}
	stakedPercentLegacyDec, err := stakedPercent.SdkLegacyDec()
	if err != nil {
		return math.LegacyDec{}, err
	}
	return stakedPercentLegacyDec, nil
}

// wrapper around emissions keeper call to get the number of blocks expected in a month
func (k Keeper) GetParamsBlocksPerMonth(ctx context.Context) (uint64, error) {
	emissionsParams, err := k.emissionsKeeper.GetParams(ctx)
	if err != nil {
		return 0, err
	}
	return emissionsParams.BlocksPerMonth, nil
}

// wrapper around emissions keeper call to set the number of blocks expected in a month
func (k Keeper) SetEmissionsParamsBlocksPerMonth(ctx context.Context, blocksPerMonth uint64) error {
	emissionsParams, err := k.emissionsKeeper.GetParams(ctx)
	if err != nil {
		return errorsmod.Wrap(err, "error getting params from emissions keeper")
	}
	emissionsParams.BlocksPerMonth = blocksPerMonth
	return k.emissionsKeeper.SetParams(ctx, emissionsParams)
}

// wrapper around emissions keeper call to get if whitelist admin
func (k Keeper) IsWhitelistAdmin(ctx context.Context, admin string) (bool, error) {
	return k.emissionsKeeper.IsWhitelistAdmin(ctx, admin)
}

// wrapper for interface compatibility for unit testing
func (k Keeper) GetPreviousRewardEmissionPerUnitStakedToken(ctx context.Context) (math.LegacyDec, error) {
	return k.PreviousRewardEmissionPerUnitStakedToken.Get(ctx)
}

// wrapper for interface compatibility for unit testing
func (k Keeper) GetEmissionsKeeperTotalStake(ctx context.Context) (math.Int, error) {
	return k.emissionsKeeper.GetTotalStake(ctx)
}

// wrapper for interface compatibility for unit testing
func (k Keeper) SetRewardCurrentBlockEmission(ctx context.Context, emission math.Int) error {
	return k.emissionsKeeper.SetRewardCurrentBlockEmission(ctx, emission)
}

// Getter for the number of months unlocked
// this Getter coerces values to be between 0 and 36
// rather than throwing errors for invalid values stored in the keeper
func (k Keeper) GetMonthsAlreadyUnlocked(ctx context.Context) math.Int {
	// 36 months is the maximum number of months that can be unlocked,
	// since tokens are on a three year vesting cycle
	thirtySix := math.NewInt(36)
	val, err := k.MonthsUnlocked.Get(ctx)
	if err != nil {
		return math.ZeroInt()
	}
	if val.IsNegative() {
		return math.ZeroInt()
	}
	if val.GT(thirtySix) {
		return thirtySix
	}
	return val
}
