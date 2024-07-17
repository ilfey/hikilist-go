package main

import (
	"fmt"
	"os"

	"github.com/ilfey/hikilist-go/api/router"
	"github.com/ilfey/hikilist-go/config"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/server"
	animeRepository "github.com/ilfey/hikilist-go/repositories/anime"
	tokenRepository "github.com/ilfey/hikilist-go/repositories/token"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
)

func main() {
	config.LoadEnvironment()

	config := config.NewConfig()

	db := database.NewDatabase(config.Database)

	// Create repositories.

	animeRepository := animeRepository.NewRepository(db)

	userRepository := userRepository.NewRepository(db)

	tokenRepository := tokenRepository.NewRepository(db)

	// Create services.

	animeService := animeService.NewService(animeRepository)

	userService := userService.NewService(userRepository)

	authService := authService.NewService(
		config.Auth,
		&authService.Dependencies{
			User:  userRepository,
			Token: tokenRepository,
		},
		// authRepository.NewRepository(db),
	)

	// Create router.
	router := &router.Router{
		AnimeService: animeService,
		AuthService:  authService,
		UserService:  userService,
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
