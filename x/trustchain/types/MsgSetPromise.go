package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPromise{}

type MsgSetPromise struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  PromiseDescription string `json:"promiseDescription" yaml:"promiseDescription"`
  PromiseKeeper string `json:"promiseKeeper" yaml:"promiseKeeper"`
  Reward string `json:"reward" yaml:"reward"`
}

func NewMsgSetPromise(creator sdk.AccAddress, id string, promiseDescription string, promiseKeeper string, reward string) MsgSetPromise {
  return MsgSetPromise{
    ID: id,
		Creator: creator,
    PromiseDescription: promiseDescription,
    PromiseKeeper: promiseKeeper,
    Reward: reward,
	}
}

func (msg MsgSetPromise) Route() string {
  return RouterKey
}

func (msg MsgSetPromise) Type() string {
  return "SetPromise"
}

func (msg MsgSetPromise) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetPromise) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetPromise) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}