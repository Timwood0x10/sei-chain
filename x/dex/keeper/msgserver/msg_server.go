package msgserver

import (
	"github.com/Timwood0x10/sei-chain/x/dex/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
)

type msgServer struct {
	keeper.Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper keeper.Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
