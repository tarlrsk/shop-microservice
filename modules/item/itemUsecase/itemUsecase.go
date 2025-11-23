package itemusecase

import itemrepository "github.com/tarlrsk/shop/modules/item/itemRepository"

type (
	ItemUseCaseService interface {
	}

	itemUseCase struct {
		itemRepository itemrepository.ItemRepositoryService
	}
)

func NewItemUseCase(itemRepository itemrepository.ItemRepositoryService) ItemUseCaseService {
	return &itemUseCase{itemRepository}
}
