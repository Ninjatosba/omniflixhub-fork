package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

func CustomMessageDecorator(bank *bankkeeper.BaseKeeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped: old,
			bank:    bank,
		}
	}
}

type CustomMessenger struct {
	wrapped wasmkeeper.Messenger
	bank    *bankkeeper.BaseKeeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	wasmvmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, addr sdk.AccAddress, msg string, cosmsg wasmvmtypes.CosmosMsg) error {
	_, err := m.wrapped.DispatchMsg(ctx, addr, msg, cosmsg)
	return err
}
