package wasmbinding

import (
	"fmt"
	"sync"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	onfttypes "github.com/OmniFlix/onft/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

var stargateWhitelist sync.Map

func init() {
	// Onft queries
	setWhitelistedQuery("/OmniFlix.onft.v1beta1.Query/Collection", &onfttypes.QueryCollectionResponse{})
	setWhitelistedQuery("/OmniFlix.onft.v1beta1.Query/Params", &onfttypes.QueryParamsResponse{})
	setWhitelistedQuery("/OmniFlix.onft.v1beta1.Query/OwnerONFTs", &onfttypes.QueryOwnerONFTsResponse{})
	setWhitelistedQuery("/OmniFlix.onft.v1beta1.Query/Supply", &onfttypes.QuerySupplyResponse{})

}

func GetWhitelistedQuery(queryPath string) (codec.ProtoMarshaler, error) {
	protoResponseAny, isWhitelisted := stargateWhitelist.Load(queryPath)
	if !isWhitelisted {
		return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("'%s' path is not allowed from the contract", queryPath)}
	}
	protoResponseType, ok := protoResponseAny.(codec.ProtoMarshaler)
	if !ok {
		return nil, wasmvmtypes.Unknown{}
	}
	return protoResponseType, nil
}

func setWhitelistedQuery(queryPath string, protoType codec.ProtoMarshaler) {
	stargateWhitelist.Store(queryPath, protoType)
}

func GetStargateWhitelistedPaths() (keys []string) {
	// Iterate over the map and collect the keys
	stargateWhitelist.Range(func(key, value interface{}) bool {
		keyStr, ok := key.(string)
		if !ok {
			panic("key is not a string")
		}
		keys = append(keys, keyStr)
		return true
	})

	return keys
}
 