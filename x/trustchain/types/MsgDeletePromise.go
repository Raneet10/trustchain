package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeletePromise{}

type MsgDeletePromise struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeletePromise(id string, creator sdk.AccAddress) MsgDeletePromise {
  return MsgDeletePromise{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeletePromise) Route() string {
  return RouterKey
}

func (msg MsgDeletePromise) Type() string {
  return "DeletePromise"
}

func (msg MsgDeletePromise) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeletePromise) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeletePromise) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}