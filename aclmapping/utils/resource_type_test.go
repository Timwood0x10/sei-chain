package utils_test

import (
	"fmt"
	"testing"

	aclutils "github.com/Timwood0x10/sei-chain/aclmapping/utils"
	sdkacltypes "github.com/cosmos/cosmos-sdk/types/accesscontrol"
)

func TestAllResourcesInTree(t *testing.T) {
	storeKeyToResourceMap := aclutils.StoreKeyToResourceTypePrefixMap
	resourceTree := sdkacltypes.ResourceTree

	storeKeyAllResourceTypes := make(map[sdkacltypes.ResourceType]bool)
	for _, resourceTypeToPrefix := range storeKeyToResourceMap {
		for resourceType := range resourceTypeToPrefix {
			storeKeyAllResourceTypes[resourceType] = true
		}
	}

	for resourceType := range resourceTree {
		if _, ok := storeKeyAllResourceTypes[resourceType]; !ok {
			panic(fmt.Sprintf("Missing resourceType=%s in the storekey to resource type prefix mapping", resourceType))
		}
	}

}
