package middlewareusecase

import middlewarerepository "github.com/tarlrsk/shop/modules/middleware/middlewareRepository"

type (
	MiddlewareUseCaseService interface{}

	middlewareUseCase struct {
		middlewareRepository middlewarerepository.MiddlewareRepositoryHandler
	}
)

func NewMiddlewareUseCase(
	middlewareRepository middlewarerepository.MiddlewareRepositoryHandler,
) MiddlewareUseCaseService {
	return &middlewareUseCase{
		middlewareRepository: middlewareRepository,
	}
}
