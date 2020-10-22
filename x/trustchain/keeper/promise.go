package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/zeno-bg/trustchain/x/trustchain/types"
)

// CreatePromise creates a promise
func (k Keeper) CreatePromise(ctx sdk.Context, promise types.Promise) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PromisePrefix + promise.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(promise)
	store.Set(key, value)
}

// GetPromise returns the promise information
func (k Keeper) GetPromise(ctx sdk.Context, key string) (types.Promise, error) {
	store := ctx.KVStore(k.storeKey)
	var promise types.Promise
	byteKey := []byte(types.PromisePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &promise)
	if err != nil {
		return promise, err
	}
	return promise, nil
}

// SetPromise sets a promise
// func (k Keeper) SetPromise(ctx sdk.Context, promise types.Promise) {
//     promiseKey := promise.ID
//     store := ctx.KVStore(k.storeKey)
//     bz := k.cdc.MustMarshalBinaryLengthPrefixed(promise)
//     key := []byte(types.PromisePrefix + promiseKey)
//     store.Set(key, bz)
// }

// // DeletePromise deletes a promise
// func (k Keeper) DeletePromise(ctx sdk.Context, key string) {
//     store := ctx.KVStore(k.storeKey)
//     store.Delete([]byte(types.PromisePrefix + key))
// }

//
// Functions used by querier
//

func listPromise(ctx sdk.Context, k Keeper) ([]byte, error) {
	var promiseList []types.Promise
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PromisePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var promise types.Promise
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &promise)
		promiseList = append(promiseList, promise)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, promiseList)
	return res, nil
}

func getPromise(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	promise, err := k.GetPromise(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, promise)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetPromiseOwner(ctx sdk.Context, key string) sdk.AccAddress {
	promise, err := k.GetPromise(ctx, key)
	if err != nil {
		return nil
	}
	return promise.Creator
}

// Check if the key exists in the store
func (k Keeper) PromiseExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PromisePrefix + key))
}
