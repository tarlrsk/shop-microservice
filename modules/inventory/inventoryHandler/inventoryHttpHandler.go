package inventoryhandler

import (
	"github.com/tarlrsk/shop/config"
	itemusecase "github.com/tarlrsk/shop/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface {
	}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUseCase itemusecase.ItemUseCaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUseCase itemusecase.ItemUseCaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		cfg:         cfg,
		itemUseCase: itemUseCase,
	}
}
