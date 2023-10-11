package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

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

