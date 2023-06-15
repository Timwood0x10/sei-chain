package query

import (
	"context"

	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k KeeperWrapper) AssetList(c context.Context, req *types.QueryAssetListRequest) (*types.QueryAssetListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	allAssetMetadata := k.GetAllAssetMetadata(ctx)

	return &types.QueryAssetListResponse{AssetList: allAssetMetadata}, nil
}

func (k KeeperWrapper) AssetMetadata(c context.Context, req *types.QueryAssetMetadataRequest) (*types.QueryAssetMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	assetMetadata, found := k.GetAssetMetadataByDenom(ctx, req.Denom)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryAssetMetadataResponse{Metadata: &assetMetadata}, nil
}
