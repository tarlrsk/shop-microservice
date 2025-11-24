package server

import (
	inventoryhandler "github.com/tarlrsk/shop/modules/inventory/inventoryHandler"
	inventoryrepository "github.com/tarlrsk/shop/modules/inventory/inventoryRepository"
	inventoryusecase "github.com/tarlrsk/shop/modules/inventory/inventoryUseCase"
)

func (s *server) inventoryService() {
	repo := inventoryrepository.NewInventoryRepository(s.db)
	useCase := inventoryusecase.NewInventoryUseCase(repo)
	httpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, useCase)
	grpcHandler := inventoryhandler.NewInventoryGrpcHandler(useCase)
	queueHandler := inventoryhandler.NewInventoryQueueHandler(s.cfg, useCase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health Check
	inventory.GET("", s.healthCheckService)
}
