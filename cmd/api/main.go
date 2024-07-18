package main

import (
	"fmt"
	"os"

	"github.com/ilfey/hikilist-go/api/router"
	"github.com/ilfey/hikilist-go/config"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/server"
	animeRepository "github.com/ilfey/hikilist-go/repositories/anime"
	tokenRepository "github.com/ilfey/hikilist-go/repositories/token"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	userActionRepository "github.com/ilfey/hikilist-go/repositories/user_action"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
	userActionService "github.com/ilfey/hikilist-go/services/user_action"
)

func main() {
	logger.SetLevel(logger.LevelDebug)

	config.LoadEnvironment()

	config := config.NewConfig()

	db := database.NewDatabase(config.Database)

	// Create repositories.

	animeRepository := animeRepository.New(db)

	userRepository := userRepository.New(db)

	userActionRepository := userActionRepository.New(db)

	tokenRepository := tokenRepository.New(db)

	// Create services.

	animeService := animeService.New(animeRepository)

	userService := userService.New(userRepository)

	authService := authService.New(
		config.Auth,
		&authService.Dependencies{
			User:  userRepository,
			Token: tokenRepository,
		},
		// authRepository.NewRepository(db),
	)

	userActionService := userActionService.New(userActionRepository)

	// Create router.
	router := &router.Router{
		AnimeService: animeService,
		AuthService:  authService,
		UserService:  userService,
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
