package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAdmin{}

func NewMsgCreateAdmin(
	creator string,
	address string,

) *MsgCreateAdmin {
	return &MsgCreateAdmin{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateAdmin) Route() string {
	return RouterKey
}

func (msg *MsgCreateAdmin) Type() string {
	return "CreateAdmin"
}

func (msg *MsgCreateAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAdmin{}

func NewMsgDeleteAdmin(
	creator string,
	address string,

) *MsgDeleteAdmin {
	return &MsgDeleteAdmin{
		Creator: creator,
		Address: address,
	}
}
func (msg *MsgDeleteAdmin) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAdmin) Type() string {
	return "DeleteAdmin"
}

func (msg *MsgDeleteAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
