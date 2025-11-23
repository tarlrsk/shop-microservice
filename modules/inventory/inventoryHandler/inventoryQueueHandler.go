package inventoryhandler

import (
	"github.com/tarlrsk/shop/config"
	inventoryusecase "github.com/tarlrsk/shop/modules/inventory/inventoryUseCase"
)

type (
	InventoryQueueHandlerService interface {
	}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUseCase inventoryusecase.InventoryUseCaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUseCase inventoryusecase.InventoryUseCaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{
		cfg:              cfg,
		inventoryUseCase: inventoryUseCase,
	}
}
