package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSuperAdmin{}

func NewMsgCreateSuperAdmin(creator string, address string) *MsgCreateSuperAdmin {
	return &MsgCreateSuperAdmin{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateSuperAdmin) Route() string {
	return RouterKey
}

func (msg *MsgCreateSuperAdmin) Type() string {
	return "CreateSuperAdmin"
}

func (msg *MsgCreateSuperAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSuperAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSuperAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSuperAdmin{}

func NewMsgUpdateSuperAdmin(creator string, address string) *MsgUpdateSuperAdmin {
	return &MsgUpdateSuperAdmin{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateSuperAdmin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSuperAdmin) Type() string {
	return "UpdateSuperAdmin"
}

func (msg *MsgUpdateSuperAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSuperAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSuperAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
