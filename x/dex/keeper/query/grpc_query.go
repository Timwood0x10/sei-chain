package query

import (
	"github.com/Timwood0x10/sei-chain/x/dex/types"
)

var _ types.QueryServer = KeeperWrapper{}
