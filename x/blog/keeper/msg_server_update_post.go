package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rohitroyrr8/blog/x/blog/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	post := types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
		Id:      msg.Id,
	}

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("ket %d does not exist", msg.Id))
	}

	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPost(ctx, post)
	return &types.MsgUpdatePostResponse{}, nil
}
