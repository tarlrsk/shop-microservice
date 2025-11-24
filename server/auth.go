package server

import (
	authhandler "github.com/tarlrsk/shop/modules/auth/authHandler"
	authrepository "github.com/tarlrsk/shop/modules/auth/authRepository"
	authusecase "github.com/tarlrsk/shop/modules/auth/authUseCase"
)

func (s *server) authService() {
	repo := authrepository.NewAuthRepository(s.db)
	useCase := authusecase.NewAuthUseCase(repo)
	httpHandler := authhandler.NewAuthHttpHandler(s.cfg, useCase)
	grpcHandler := authhandler.NewAuthGrpcHandler(useCase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}
