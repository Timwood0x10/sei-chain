package keeper

import (
	"github.com/Timwood0x10/sei-chain/x/epoch/types"
)

var _ types.QueryServer = Keeper{}
