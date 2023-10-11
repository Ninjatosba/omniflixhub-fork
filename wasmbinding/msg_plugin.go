package wasmbinding

import (
	"encoding/json"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	bindings "github.com/OmniFlix/omniflixhub/wasmbinding/bindings"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	onftkeeper "github.com/OmniFlix/onft/keeper"
)

func CustomMessageDecorator(bank *bankkeeper.BaseKeeper, onftKeeper *onftkeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:    old,
			bank:       bank,
			onftKeeper: onftKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped    wasmkeeper.Messenger
	bank       *bankkeeper.BaseKeeper
	onftKeeper *onftkeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {

	// Print something to the console
	ctx.Logger().Info("Custom message received", "msg", msg)
	if msg.Custom != nil {
		// only handle the happy path where this is really creating / minting / swapping ...
		// leave everything else for the wrapped version
		var contractMsg bindings.OmniflixMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			// Print something to the console
			ctx.Logger().Info("Custom message failed to unmarshal", "err", err)
		}
		// log the message
		ctx.Logger().Info("Custom message parsed", "msg", contractMsg)
		if contractMsg.TestBinding != nil {
			// Print something to the console
			ctx.Logger().Info("Custom message received", "test", contractMsg.TestBinding)
		}
		if contractMsg.CreateDenom != nil {
			// Print something to the console
			ctx.Logger().Info("Custom message received", "create_denom", contractMsg.CreateDenom)
		}
		if contractMsg.MsgCreateDenom != nil {
			// Print something to the console
			ctx.Logger().Info("Custom message received", "msg_create_denom", contractMsg.MsgCreateDenom)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}
