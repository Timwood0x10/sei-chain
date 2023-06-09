package keeper_test

import (
	"testing"

	keepertest "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/testutil/nullify"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestAddGetPair(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	keeper.AddRegisteredPair(ctx, keepertest.TestContract, types.Pair{
		PriceDenom:       keepertest.TestPriceDenom,
		AssetDenom:       keepertest.TestAssetDenom,
		PriceTicksize:    &keepertest.TestTicksize,
		QuantityTicksize: &keepertest.TestTicksize,
	})
	require.ElementsMatch(t,
		nullify.Fill([]types.Pair{{
			PriceDenom:       keepertest.TestPriceDenom,
			AssetDenom:       keepertest.TestAssetDenom,
			PriceTicksize:    &keepertest.TestTicksize,
			QuantityTicksize: &keepertest.TestTicksize,
		}}),
		nullify.Fill(keeper.GetAllRegisteredPairs(ctx, keepertest.TestContract)),
	)

	pair, found := keeper.GetRegisteredPair(ctx, keepertest.TestContract, keepertest.TestPriceDenom, keepertest.TestAssetDenom)
	require.True(t, found)
	require.Equal(t, types.Pair{
		PriceDenom:       keepertest.TestPriceDenom,
		AssetDenom:       keepertest.TestAssetDenom,
		PriceTicksize:    &keepertest.TestTicksize,
		QuantityTicksize: &keepertest.TestTicksize,
	}, pair)
	hasPair := keeper.HasRegisteredPair(ctx, keepertest.TestContract, keepertest.TestPriceDenom, keepertest.TestAssetDenom)
	require.True(t, hasPair)

}
