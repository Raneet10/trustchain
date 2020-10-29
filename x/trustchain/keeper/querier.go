package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/zeno-bg/trustchain/x/trustchain/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for trustchain clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListPromise:
			return listPromise(ctx, k)
		case types.QueryListPromisesForKeeper:
			return listPromisesForKeeper(ctx, path[1], k)
		case types.QueryGetPromise:
			return getPromise(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown trustchain query endpoint")
		}
	}
}
