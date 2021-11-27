package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendToSystem{}

func NewMsgSendToSystem(creator string, amount uint64) *MsgSendToSystem {
	return &MsgSendToSystem{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgSendToSystem) Route() string {
	return RouterKey
}

func (msg *MsgSendToSystem) Type() string {
	return "SendToSystem"
}

func (msg *MsgSendToSystem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendToSystem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendToSystem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
