package playerhandler

import (
	"github.com/tarlrsk/shop/config"
	playerusecase "github.com/tarlrsk/shop/modules/player/playerUseCase"
)

type (
	PlayerQueueHandlerService interface {
	}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUseCase playerusecase.PlayerUseCaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUseCase playerusecase.PlayerUseCaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{
		cfg:           cfg,
		playerUseCase: playerUseCase,
	}
}
