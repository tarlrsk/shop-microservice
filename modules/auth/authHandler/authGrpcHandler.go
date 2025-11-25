package authhandler

import (
	"context"

	authPb "github.com/tarlrsk/shop/modules/auth/authPb"
	authusecase "github.com/tarlrsk/shop/modules/auth/authUseCase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUseCase authusecase.AuthUseCaseService
	}
)

func NewAuthGrpcHandler(authUseCase authusecase.AuthUseCaseService) *authGrpcHandler {
	return &authGrpcHandler{
		authUseCase: authUseCase,
	}
}

func (g *authGrpcHandler) CredentialSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return &authPb.AccessTokenSearchRes{}, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return &authPb.RolesCountRes{}, nil
}
