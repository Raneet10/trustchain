package trustchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zeno-bg/trustchain/x/trustchain/types"
	"github.com/zeno-bg/trustchain/x/trustchain/keeper"
)

// Handle a message to delete name
func handleMsgDeletePromise(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeletePromise) (*sdk.Result, error) {
	if !k.PromiseExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetPromiseOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeletePromise(ctx, msg.ID)
	return &sdk.Result{}, nil
}
