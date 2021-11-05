package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAdmin{}, "medipoint/CreateAdmin", nil)
	cdc.RegisterConcrete(&MsgDeleteAdmin{}, "medipoint/DeleteAdmin", nil)
	cdc.RegisterConcrete(&MsgCreateSuperAdmin{}, "medipoint/CreateSuperAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateSuperAdmin{}, "medipoint/UpdateSuperAdmin", nil)
	cdc.RegisterConcrete(&MsgMint{}, "medipoint/Mint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "medipoint/Burn", nil)
	cdc.RegisterConcrete(&MsgAdminTransfer{}, "medipoint/AdminTransfer", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAdmin{},
		&MsgDeleteAdmin{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSuperAdmin{},
		&MsgUpdateSuperAdmin{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMint{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAdminTransfer{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
