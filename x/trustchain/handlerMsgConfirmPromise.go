package trustchain

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zeno-bg/trustchain/x/trustchain/keeper"
	"github.com/zeno-bg/trustchain/x/trustchain/types"
)

func handleMsgConfirmPromise(ctx sdk.Context, k keeper.Keeper, msg types.MsgConfirmPromise) (*sdk.Result, error) {

	//well its not working

	promise, err := k.GetPromise(ctx, msg.PromiseId)

	if err != nil {
		return nil, err
	}

	if !promise.PromiseKeeper.Equals(msg.Creator) {
		return nil, nil //TODO print error
	}

	if !k.GetPromiseReward(ctx, msg.PromiseId).IsEqual(msg.Reward) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Incorrect Reward Amount")
	}

	if k.GetPromiseDeadline(ctx, msg.PromiseId).Before(time.Now()) {
		//Deadline exceeded error
		return nil, nil
	}

	err = k.CoinKeeper.SendCoins(ctx, msg.Creator, msg.Fulfiller, msg.Reward)

	if err != nil {
		return nil, err
	}
	k.ConfirmPromise(ctx, promise)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeConfirmPromise),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeId, msg.PromiseId),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
