package inventoryhandler

import (
	"github.com/tarlrsk/shop/config"
	inventoryusecase "github.com/tarlrsk/shop/modules/inventory/inventoryUseCase"
)

type (
	InventoryHttpHandlerService interface {
	}

	inventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUseCase inventoryusecase.InventoryUseCaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUseCase inventoryusecase.InventoryUseCaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{
		cfg:              cfg,
		inventoryUseCase: inventoryUseCase,
	}
}
