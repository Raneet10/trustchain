package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreatePromise{}

type MsgCreatePromise struct {
	ID                 string
	Creator            sdk.AccAddress `json:"creator" yaml:"creator"`
	PromiseDescription string         `json:"promiseDescription" yaml:"promiseDescription"`
	PromiseKeeper      sdk.AccAddress `json:"promiseKeeper" yaml:"promiseKeeper"`
	Reward             sdk.Coins      `json:"reward" yaml:"reward"`
}

func NewMsgCreatePromise(creator sdk.AccAddress, promiseDescription string, promiseKeeper sdk.AccAddress, reward sdk.Coins) MsgCreatePromise {
	return MsgCreatePromise{
		ID:                 uuid.New().String(),
		Creator:            creator,
		PromiseDescription: promiseDescription,
		PromiseKeeper:      promiseKeeper,
		Reward:             reward,
	}
}

func (msg MsgCreatePromise) Route() string {
	return RouterKey
}

func (msg MsgCreatePromise) Type() string {
	return "CreatePromise"
}

func (msg MsgCreatePromise) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreatePromise) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreatePromise) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.PromiseDescription == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Promise description can't be empty")
	}
	if msg.PromiseKeeper.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "promiseKeeper can't be empty")
	}
	if !msg.Reward.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "reward can't be negative nor zero")
	}
	return nil
}
