package query

import (
	"context"

	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k KeeperWrapper) GetRegisteredContract(c context.Context, req *types.QueryRegisteredContractRequest) (*types.QueryRegisteredContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	contractInfo, err := k.GetContract(ctx, req.ContractAddr)
	if err != nil {
		return nil, err
	}

	return &types.QueryRegisteredContractResponse{ContractInfo: &contractInfo}, nil
}
