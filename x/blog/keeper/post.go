package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rohitroyrr8/blog/x/blog/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))

	appendValue := k.cdc.MustMarshal(&post)
	store.Set(GetPostIdBytes(post.Id), appendValue)

	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})

	byteKey := types.KeyPrefix(types.PostCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)

}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))

	b := store.Get(GetPostIdBytes(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true

}

func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))

	b := k.cdc.MustMarshal(&post)
	store.Set(GetPostIdBytes(post.Id), b)
}

func (k Keeper) RemovePost(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	store.Delete(GetPostIdBytes(id))
}

func GetPostIdBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz

}
