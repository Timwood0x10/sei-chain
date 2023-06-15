package query

import (
	"context"

	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k KeeperWrapper) GetLatestPrice(goCtx context.Context, req *types.QueryGetLatestPriceRequest) (*types.QueryGetLatestPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	prices := k.GetAllPrices(ctx, req.ContractAddr, types.Pair{PriceDenom: req.PriceDenom, AssetDenom: req.AssetDenom})

	if len(prices) == 0 {
		return &types.QueryGetLatestPriceResponse{
			Price: &types.Price{},
		}, nil
	}

	latestPrice := prices[0]

	for _, price := range prices {
		if price.SnapshotTimestampInSeconds > latestPrice.SnapshotTimestampInSeconds {
			latestPrice = price
		}
	}

	return &types.QueryGetLatestPriceResponse{
		Price: latestPrice,
	}, nil
}
