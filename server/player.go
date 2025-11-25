package server

import (
	"log"

	playerhandler "github.com/tarlrsk/shop/modules/player/playerHandler"
	playerPb "github.com/tarlrsk/shop/modules/player/playerPb"
	playerrepository "github.com/tarlrsk/shop/modules/player/playerRepository"
	playerusecase "github.com/tarlrsk/shop/modules/player/playerUseCase"
	"github.com/tarlrsk/shop/pkg/grpccon"
)

func (s *server) playerService() {
	repo := playerrepository.NewPlayerRepository(s.db)
	useCase := playerusecase.NewPlayerUseCase(repo)
	httpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, useCase)
	grpcHandler := playerhandler.NewPlayerGrpcHandler(useCase)
	queueHandler := playerhandler.NewPlayerQueueHandler(s.cfg, useCase)

	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Starting Player gRPC server at %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("", s.healthCheckService)
}
