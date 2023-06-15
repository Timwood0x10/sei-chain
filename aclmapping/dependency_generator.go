package aclmapping

import (
	aclbankmapping "github.com/Timwood0x10/sei-chain/aclmapping/bank"
	acldexmapping "github.com/Timwood0x10/sei-chain/aclmapping/dex"
	acloraclemapping "github.com/Timwood0x10/sei-chain/aclmapping/oracle"
	acltokenfactorymapping "github.com/Timwood0x10/sei-chain/aclmapping/tokenfactory"
	aclwasmmapping "github.com/Timwood0x10/sei-chain/aclmapping/wasm"
	aclkeeper "github.com/cosmos/cosmos-sdk/x/accesscontrol/keeper"
)

type CustomDependencyGenerator struct{}

func NewCustomDependencyGenerator() CustomDependencyGenerator {
	return CustomDependencyGenerator{}
}

func (customDepGen CustomDependencyGenerator) GetCustomDependencyGenerators() aclkeeper.DependencyGeneratorMap {
	dependencyGeneratorMap := make(aclkeeper.DependencyGeneratorMap)
	wasmDependencyGenerators := aclwasmmapping.NewWasmDependencyGenerator()

	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acldexmapping.GetDexDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(aclbankmapping.GetBankDepedencyGenerator())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acltokenfactorymapping.GetTokenFactoryDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(wasmDependencyGenerators.GetWasmDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acloraclemapping.GetOracleDependencyGenerator())

	return dependencyGeneratorMap
}
