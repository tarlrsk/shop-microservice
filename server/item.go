package server

import (
	itemhandler "github.com/tarlrsk/shop/modules/item/itemHandler"
	itemrepository "github.com/tarlrsk/shop/modules/item/itemRepository"
	itemusecase "github.com/tarlrsk/shop/modules/item/itemUseCase"
)

func (s *server) itemService() {
	repo := itemrepository.NewItemRepository(s.db)
	useCase := itemusecase.NewItemUseCase(repo)
	httpHandler := itemhandler.NewItemHttpHandler(s.cfg, useCase)
	grpcHandler := itemhandler.NewItemGrpcHandler(useCase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Check
	item.GET("", s.healthCheckService)
}
