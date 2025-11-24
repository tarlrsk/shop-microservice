package authhandler

import (
	"github.com/tarlrsk/shop/config"
	authusecase "github.com/tarlrsk/shop/modules/auth/authUseCase"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUseCase authusecase.AuthUseCaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUseCase authusecase.AuthUseCaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		cfg:         cfg,
		authUseCase: authUseCase,
	}
}
