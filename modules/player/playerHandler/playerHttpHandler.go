package playerhandler

import (
	"github.com/tarlrsk/shop/config"
	playerusecase "github.com/tarlrsk/shop/modules/player/playerUseCase"
)

type (
	PlayerHttpHandlerService interface {
	}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUseCase playerusecase.PlayerUseCaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUseCase playerusecase.PlayerUseCaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{
		cfg:           cfg,
		playerUseCase: playerUseCase,
	}
}
