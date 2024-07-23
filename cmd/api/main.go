package main

import (
	"fmt"
	"os"

	"github.com/ilfey/hikilist-go/api/router"
	"github.com/ilfey/hikilist-go/config"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/server"

	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	collectionService "github.com/ilfey/hikilist-go/services/collection"
	userService "github.com/ilfey/hikilist-go/services/user"
	userActionService "github.com/ilfey/hikilist-go/services/user_action"
)

func main() {
	logger.SetLevel(logger.LevelDebug)

	config.LoadEnvironment()

	config := config.NewConfig()

	db := database.NewDatabase(config.Database)

	// Create services.

	animeService := animeService.New(db)

	userService := userService.New(db)

	authService := authService.New(
		config.Auth,
		db,
	)

	collectionService := collectionService.New(db)

	userActionService := userActionService.New(db)

	// Create router.
	router := &router.Router{
		AnimeService:      animeService,
		AuthService:       authService,
		CollectionService: collectionService,
		UserService:       userService,
		UserActionService: userActionService,
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
