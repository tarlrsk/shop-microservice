package playerusecase

import playerrepository "github.com/tarlrsk/shop/modules/player/playerRepository"

type (
	PlayerUseCaseService interface {
	}

	playerUseCase struct {
		playerRepository playerrepository.PlayerRepositoryService
	}
)

func NewPlayerUseCase(playerRepository playerrepository.PlayerRepositoryService) PlayerUseCaseService {
	return &playerUseCase{playerRepository: playerRepository}
}
