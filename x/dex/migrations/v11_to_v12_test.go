package migrations_test

import (
	"testing"

	keepertest "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/migrations"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMigrate11to12(t *testing.T) {
	dexkeeper, ctx := keepertest.DexKeeper(t)
	// write old params
	defaultParams := types.Params{
		PriceSnapshotRetention: 1,
		SudoCallGasPrice:       sdk.OneDec(),
		BeginBlockGasLimit:     1,
		EndBlockGasLimit:       1,
		DefaultGasPerOrder:     1,
		DefaultGasPerCancel:    1,
	}
	dexkeeper.SetParams(ctx, defaultParams)

	// migrate to default params
	err := migrations.V11ToV12(ctx, *dexkeeper)
	require.NoError(t, err)
	params := dexkeeper.GetParams(ctx)
	require.Equal(t, params, types.DefaultParams())
}
