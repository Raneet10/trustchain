package trustchain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zeno-bg/trustchain/x/trustchain/keeper"
	"github.com/zeno-bg/trustchain/x/trustchain/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgCreatePromise:
			return handleMsgCreatePromise(ctx, k, msg)
		case types.MsgConfirmPromise:
			return handleMsgConfirmPromise(ctx, k, msg)
		// case types.MsgSetPromise:
		//     return handleMsgSetPromise(ctx, k, msg)
		// case types.MsgDeletePromise:
		//     return handleMsgDeletePromise(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
