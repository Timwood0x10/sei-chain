package query

import (
	"context"

	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k KeeperWrapper) GetRegisteredPairs(c context.Context, req *types.QueryRegisteredPairsRequest) (*types.QueryRegisteredPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	registeredPairs := k.GetAllRegisteredPairs(ctx, req.ContractAddr)

	return &types.QueryRegisteredPairsResponse{Pairs: registeredPairs}, nil
}
