package server

import (
	"log"

	itemhandler "github.com/tarlrsk/shop/modules/item/itemHandler"
	"github.com/tarlrsk/shop/modules/item/itemPb"
	itemrepository "github.com/tarlrsk/shop/modules/item/itemRepository"
	itemusecase "github.com/tarlrsk/shop/modules/item/itemUseCase"
	"github.com/tarlrsk/shop/pkg/grpccon"
)

func (s *server) itemService() {
	repo := itemrepository.NewItemRepository(s.db)
	useCase := itemusecase.NewItemUseCase(repo)
	httpHandler := itemhandler.NewItemHttpHandler(s.cfg, useCase)
	grpcHandler := itemhandler.NewItemGrpcHandler(useCase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Starting Item gRPC server at %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Check
	item.GET("", s.healthCheckService)
}
