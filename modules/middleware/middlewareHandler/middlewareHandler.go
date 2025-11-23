package middlewarehandler

import (
	"github.com/tarlrsk/shop/config"
	middlewareusecase "github.com/tarlrsk/shop/modules/middleware/middlewareUseCase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUseCase middlewareusecase.MiddlewareUseCaseService
	}
)

func NewMiddlewareHandler(
	cfg *config.Config,
	middlewareUseCase middlewareusecase.MiddlewareUseCaseService,
) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUseCase: middlewareUseCase,
	}
}
