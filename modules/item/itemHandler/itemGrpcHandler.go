package itemhandler

import (
	itemusecase "github.com/tarlrsk/shop/modules/item/itemUseCase"
)

type (
	itemGrpcHandler struct {
		itemUseCase itemusecase.ItemUseCaseService
	}
)

func NewItemGrpcHandler(itemUseCase itemusecase.ItemUseCaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUseCase: itemUseCase,
	}
}
