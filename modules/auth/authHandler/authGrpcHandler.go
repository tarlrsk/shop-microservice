package authhandler

import (
	authusecase "github.com/tarlrsk/shop/modules/auth/authUseCase"
)

type (
	authGrpcHandler struct {
		authUseCase authusecase.AuthUseCaseService
	}
)

func NewAuthGrpcHandler(authUseCase authusecase.AuthUseCaseService) *authGrpcHandler {
	return &authGrpcHandler{
		authUseCase: authUseCase,
	}
}
