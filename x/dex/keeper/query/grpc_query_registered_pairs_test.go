package query_test

import (
	"testing"

	keepertest "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/keeper/query"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestRegisteredPairsQuery(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wrapper := query.KeeperWrapper{Keeper: keeper}
	wctx := sdk.WrapSDKContext(ctx)
	expectedPair := types.Pair{
		PriceDenom:       keepertest.TestPriceDenom,
		AssetDenom:       keepertest.TestAssetDenom,
		PriceTicksize:    &keepertest.TestTicksize,
		QuantityTicksize: &keepertest.TestTicksize,
	}
	keeper.AddRegisteredPair(ctx, keepertest.TestContract, expectedPair)

	var expectedRegisteredPairs []types.Pair
	expectedRegisteredPairs = append(expectedRegisteredPairs, expectedPair)

	request := types.QueryRegisteredPairsRequest{
		ContractAddr: keepertest.TestContract,
	}
	expectedResponse := types.QueryRegisteredPairsResponse{
		Pairs: expectedRegisteredPairs,
	}
	t.Run("Registered Pairs query", func(t *testing.T) {
		response, err := wrapper.GetRegisteredPairs(wctx, &request)
		require.NoError(t, err)
		require.Equal(t, expectedResponse, *response)
	})
}
