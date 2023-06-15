package bindings

import "github.com/Timwood0x10/sei-chain/x/epoch/types"

type SeiEpochQuery struct {
	// queries the current Epoch
	Epoch *types.QueryEpochRequest `json:"epoch,omitempty"`
}
