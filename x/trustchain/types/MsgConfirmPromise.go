package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgConfirmPromise{}

type MsgConfirmPromise struct {
	PromiseId string
	Reward    sdk.Coins
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	Fulfiller sdk.AccAddress `json:"fulfiller" yaml:"fulfiller"`
}

func NewMsgConfirmPromise(creator sdk.AccAddress, fulfiller sdk.AccAddress, reward sdk.Coins, promiseId string) MsgConfirmPromise {
	return MsgConfirmPromise{
		PromiseId: uuid.New().String(),
		Reward:    reward,
		Creator:   creator,
		Fulfiller: fulfiller,
	}
}

func (msg MsgConfirmPromise) Route() string {
	return RouterKey
}

func (msg MsgConfirmPromise) Type() string {
	return "ConfirmPromise"
}

func (msg MsgConfirmPromise) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgConfirmPromise) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgConfirmPromise) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	if msg.Fulfiller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "promise fulfiller can't be empty")
	}

	if !msg.Reward.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}
