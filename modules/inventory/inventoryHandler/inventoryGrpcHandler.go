package inventoryhandler

import (
	inventoryusecase "github.com/tarlrsk/shop/modules/inventory/inventoryUseCase"
)

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryusecase.InventoryUseCaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryusecase.InventoryUseCaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
