package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
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
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}
 