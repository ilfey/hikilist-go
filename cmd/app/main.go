package main

import (
	"github.com/ilfey/hikilist-go/api/router"
	"github.com/ilfey/hikilist-go/config"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/server"
	animeRepository "github.com/ilfey/hikilist-go/repositories/anime"
	authRepository "github.com/ilfey/hikilist-go/repositories/auth"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
)

func main() {
	config.LoadEnvironment()

	config := config.NewConfig()

	db := database.NewDatabase(config.Database)

	animeService := animeService.NewService(
		animeRepository.NewRepository(db),
	)

	userService := userService.NewService(
		userRepository.NewRepository(db),
	)

	authService := authService.NewService(
		config.Auth,
		authRepository.NewRepository(db),
	)

	router := &router.Router{
		AnimeService: animeService,
		AuthService:  authService,
		UserService:  userService,
	}

	srv := server.NewServer(config.Server, router)

	srv.Run()
}
