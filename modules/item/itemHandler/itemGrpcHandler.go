package itemhandler

import (
	"context"

	itemPb "github.com/tarlrsk/shop/modules/item/itemPb"
	itemusecase "github.com/tarlrsk/shop/modules/item/itemUseCase"
)

type (
	itemGrpcHandler struct {
		itemPb.UnimplementedItemGrpcServiceServer
		itemUseCase itemusecase.ItemUseCaseService
	}
)

func NewItemGrpcHandler(itemUseCase itemusecase.ItemUseCaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUseCase: itemUseCase,
	}
}

func (g *itemGrpcHandler) FindItemInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return nil, nil
}
