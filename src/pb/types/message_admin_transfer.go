package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAdminTransfer{}

func NewMsgAdminTransfer(creator string, address string, amount uint64) *MsgAdminTransfer {
	return &MsgAdminTransfer{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgAdminTransfer) Route() string {
	return RouterKey
}

func (msg *MsgAdminTransfer) Type() string {
	return "AdminTransfer"
}

func (msg *MsgAdminTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAdminTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAdminTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
