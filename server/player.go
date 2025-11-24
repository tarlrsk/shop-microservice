package server

import (
	playerhandler "github.com/tarlrsk/shop/modules/player/playerHandler"
	playerrepository "github.com/tarlrsk/shop/modules/player/playerRepository"
	playerusecase "github.com/tarlrsk/shop/modules/player/playerUseCase"
)

func (s *server) playerService() {
	repo := playerrepository.NewPlayerRepository(s.db)
	useCase := playerusecase.NewPlayerUseCase(repo)
	httpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, useCase)
	grpcHandler := playerhandler.NewPlayerGrpcHandler(useCase)
	queueHandler := playerhandler.NewPlayerQueueHandler(s.cfg, useCase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("", s.healthCheckService)
}
