package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/CosmWasm/wasmd/x/wasm"
	dexcache "github.com/Timwood0x10/sei-chain/x/dex/cache"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	epochkeeper "github.com/Timwood0x10/sei-chain/x/epoch/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		Cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		Paramstore    paramtypes.Subspace
		AccountKeeper authkeeper.AccountKeeper
		EpochKeeper   epochkeeper.Keeper
		BankKeeper    bankkeeper.Keeper
		WasmKeeper    wasm.Keeper
		MemState      *dexcache.MemState
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	epochKeeper epochkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	return &Keeper{
		Cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		Paramstore:    ps,
		EpochKeeper:   epochKeeper,
		BankKeeper:    bankKeeper,
		AccountKeeper: accountKeeper,
		MemState:      dexcache.NewMemState(memKey),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetStoreKey() sdk.StoreKey {
	return k.storeKey
}

func (k Keeper) GetMemStoreKey() sdk.StoreKey {
	return k.memKey
}

func (k *Keeper) SetWasmKeeper(wasmKeeper *wasm.Keeper) {
	k.WasmKeeper = *wasmKeeper
}

func (k Keeper) CreateModuleAccount(ctx sdk.Context) {
	moduleAcc := authtypes.NewEmptyModuleAccount(types.ModuleName)
	k.AccountKeeper.SetModuleAccount(ctx, moduleAcc)
}
