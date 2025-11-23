package main

import (
	"context"
	"log"
	"os"

	"github.com/tarlrsk/shop/config"
	"github.com/tarlrsk/shop/pkg/database"
	"github.com/tarlrsk/shop/server"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	db := database.DbConn(ctx, cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, cfg, db)
}
