package playerhandler

import playerusecase "github.com/tarlrsk/shop/modules/player/playerUseCase"

type (
	playerGrpcHandler struct {
		playerUseCase playerusecase.PlayerUseCaseService
	}
)

func NewPlayerGrpcHandler(playerUseCase playerusecase.PlayerUseCaseService) *playerGrpcHandler {
	return &playerGrpcHandler{
		playerUseCase: playerUseCase,
	}
}
