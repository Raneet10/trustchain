package trustchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/zeno-bg/trustchain/x/trustchain/keeper"
	"github.com/zeno-bg/trustchain/x/trustchain/types"
)

func handleMsgCreatePromise(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePromise) (*sdk.Result, error) {
	var promise = types.Promise{
		Creator:            msg.Creator,
		ID:                 msg.ID,
		PromiseDescription: msg.PromiseDescription,
		PromiseKeeper:      msg.PromiseKeeper,
		Reward:             msg.Reward,
	}

	moduleAccount := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	sdkError := k.CoinKeeper.SendCoins(ctx, promise.Creator, moduleAccount, promise.Reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.CreatePromise(ctx, promise)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreatePromise),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributePromiseDescription, msg.PromiseDescription),
			sdk.NewAttribute(types.AttributePromiseKeeper, msg.PromiseKeeper.String()),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
