package inventoryusecase

import inventoryrepository "github.com/tarlrsk/shop/modules/inventory/inventoryRepository"

type (
	InventoryUseCaseService interface {
	}

	inventoryUseCase struct {
		inventoryRepository inventoryrepository.InventoryRepositoryService
	}
)

func NewInventoryUseCase(inventoryRepository inventoryrepository.InventoryRepositoryService) InventoryUseCaseService {
	return &inventoryUseCase{
		inventoryRepository: inventoryRepository,
	}
}
