package playerhandler

import (
	"context"

	playerPb "github.com/tarlrsk/shop/modules/player/playerPb"
	playerusecase "github.com/tarlrsk/shop/modules/player/playerUseCase"
)

type (
	playerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerUseCase playerusecase.PlayerUseCaseService
	}
)

func NewPlayerGrpcHandler(playerUseCase playerusecase.PlayerUseCaseService) *playerGrpcHandler {
	return &playerGrpcHandler{
		playerUseCase: playerUseCase,
	}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return &playerPb.PlayerProfile{}, nil
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return &playerPb.PlayerProfile{}, nil
}

func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return &playerPb.GetPlayerSavingAccountRes{}, nil
}
