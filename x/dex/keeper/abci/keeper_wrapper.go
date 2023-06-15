package abci

import (
	"github.com/Timwood0x10/sei-chain/x/dex/keeper"
)

type KeeperWrapper struct {
	*keeper.Keeper
}
