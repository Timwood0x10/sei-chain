package multi

import (
	"github.com/Timwood0x10/sei-chain/store/whitelist/cachemulti"
	"github.com/Timwood0x10/sei-chain/store/whitelist/kv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

type Store struct {
	storetypes.MultiStore

	storeKeyToWriteWhitelist map[string][]string
}

func NewStore(parent storetypes.MultiStore, storeKeyToWriteWhitelist map[string][]string) storetypes.MultiStore {
	return &Store{
		MultiStore:               parent,
		storeKeyToWriteWhitelist: storeKeyToWriteWhitelist,
	}
}

func (cms Store) CacheMultiStore() storetypes.CacheMultiStore {
	return cachemulti.NewStore(cms.MultiStore.CacheMultiStore(), cms.storeKeyToWriteWhitelist)
}

func (cms Store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	rawKVStore := cms.MultiStore.GetKVStore(key)
	if writeWhitelist, ok := cms.storeKeyToWriteWhitelist[key.Name()]; ok {
		return kv.NewStore(rawKVStore, writeWhitelist)
	}
	// whitelist nothing
	return kv.NewStore(rawKVStore, []string{})
}
