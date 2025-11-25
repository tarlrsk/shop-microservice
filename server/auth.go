package server

import (
	"log"

	authhandler "github.com/tarlrsk/shop/modules/auth/authHandler"
	authPb "github.com/tarlrsk/shop/modules/auth/authPb"
	authrepository "github.com/tarlrsk/shop/modules/auth/authRepository"
	authusecase "github.com/tarlrsk/shop/modules/auth/authUseCase"
	"github.com/tarlrsk/shop/pkg/grpccon"
)

func (s *server) authService() {
	repo := authrepository.NewAuthRepository(s.db)
	useCase := authusecase.NewAuthUseCase(repo)
	httpHandler := authhandler.NewAuthHttpHandler(s.cfg, useCase)
	grpcHandler := authhandler.NewAuthGrpcHandler(useCase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Starting Auth gRPC server at %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}
