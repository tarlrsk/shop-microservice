package inventoryhandler

import (
	itemusecase "github.com/tarlrsk/shop/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemusecase.ItemUseCaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemusecase.ItemUseCaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}
