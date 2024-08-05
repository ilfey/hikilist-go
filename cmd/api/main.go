package main

import (
	"fmt"
	"os"

	"github.com/ilfey/hikilist-go/api/router"
	"github.com/ilfey/hikilist-go/config"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/server"

	authService "github.com/ilfey/hikilist-go/services/auth"
)

func main() {
	logger.SetLevel(logger.LevelTrace)

	config.LoadEnvironment()

	config := config.New()

	database.New(config.Database)

	// Create services.

	authService := authService.New(
		config.Auth,
	)

	// Create router.
	router := &router.Router{
		AuthService: authService,
	}

	// Create server.
	srv := server.NewServer(config.Server, router)

	// Run server.
	err := srv.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
