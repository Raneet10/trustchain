package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgConfirmPromise{}

type MsgConfirmPromise struct {
	PromiseId string
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgConfirmPromise(creator sdk.AccAddress, promiseId string) MsgConfirmPromise {
	return MsgConfirmPromise{
		PromiseId: uuid.New().String(),
		Creator:   creator,
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
	return nil
}
