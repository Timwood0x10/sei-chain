package dex

import (
	"github.com/Timwood0x10/sei-chain/x/dex/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func HandleAddAssetMetadataProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddAssetMetadataProposal) error {
	for _, asset := range p.AssetList {
		k.SetAssetMetadata(ctx, asset)
	}
	return nil
}
