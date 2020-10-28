package trustchain

// import (
//     sdk "github.com/cosmos/cosmos-sdk/types"
//     sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

//     "github.com/zeno-bg/trustchain/x/trustchain/types"
//     "github.com/zeno-bg/trustchain/x/trustchain/keeper"
// )

// func handleMsgSetPromise(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetPromise) (*sdk.Result, error) {
//     var promise = types.Promise{
//         Creator: msg.Creator,
//         ID:      msg.ID,
//         PromiseDescription: msg.PromiseDescription,
//         PromiseKeeper: msg.PromiseKeeper,
//         Reward: msg.Reward,
//     }
//     if !msg.Creator.Equals(k.GetPromiseOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
//         return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
//     }

//     k.SetPromise(ctx, promise)

//     return &sdk.Result{Events: ctx.EventManager().Events()}, nil
// }
