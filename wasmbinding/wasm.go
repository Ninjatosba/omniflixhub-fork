package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	onftkeeper "github.com/OmniFlix/onft/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	onft *onftkeeper.Keeper,
) []wasmkeeper.Option {
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, onft),
	)
	return []wasm.Option{
		messengerDecoratorOpt,
	}
}

func RegisterStargateQueries(queryRouter baseapp.GRPCQueryRouter, codec codec.Codec) []wasmkeeper.Option {
	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Stargate: StargateQuerier(queryRouter, codec),
	})

	return []wasm.Option{
		queryPluginOpt,
	}
}
