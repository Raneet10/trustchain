package trustchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

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

	sdkError := k.CoinKeeper.SendCoins(ctx, promise.Creator, moduleAccount, promise.Reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.CreatePromise(ctx, promise)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
