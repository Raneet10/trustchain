package trustchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/zeno-bg/trustchain/x/trustchain/keeper"
	"github.com/zeno-bg/trustchain/x/trustchain/types"
)

func handleMsgConfirmPromise(ctx sdk.Context, k keeper.Keeper, msg types.MsgConfirmPromise) (*sdk.Result, error) {

	promise, err := k.GetPromise(ctx, msg.PromiseId)

	if err != nil {
		return nil, err
	}

	if !promise.PromiseKeeper.Equals(msg.Creator) {
		return nil, nil //TODO print error
	}

	moduleAccount := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	sdkError := k.CoinKeeper.SendCoins(ctx, promise.Creator, moduleAccount, promise.Reward)

	if sdkError != nil {
		return nil, sdkError
	}

	promise.Confirmed = true

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
