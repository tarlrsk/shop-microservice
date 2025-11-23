package authusecase

import authrepository "github.com/tarlrsk/shop/modules/auth/authRepository"

type (
	AuthUseCaseService interface{}

	authUseCase struct {
		authRepository authrepository.AuthRepositoryService
	}
)

func NewAuthUseCase(authRepo authrepository.AuthRepositoryService) AuthUseCaseService {
	return &authUseCase{
		authRepository: authRepo,
	}
}
