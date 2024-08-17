package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/server"
	"github.com/ilfey/hikilist-go/pkg/api/router"
	"github.com/ilfey/hikilist-go/pkg/config"
	"github.com/ilfey/hikilist-go/pkg/database"
	"github.com/ilfey/hikilist-go/pkg/repositories"

	"github.com/ilfey/hikilist-go/pkg/services"
)

func main() {
	logger.SetLevel(logger.LevelTrace)

	config.LoadEnvironment()

	config := config.New()

	db := database.New(config.Database)

	// Create repositories.

	actionRepo := repositories.NewAction(db)
	animeRepo := repositories.NewAnime(db)
	animeCollectionRepo := repositories.NewAnimeCollection(db)
	collectionRepo := repositories.NewCollection(db, actionRepo)
	userRepo := repositories.NewUser(db, actionRepo)
	tokenRepo := repositories.NewToken(db)

	// Create services.

	action := services.NewAction(actionRepo)

	anime := services.NewAnime(animeRepo)

	auth := services.NewAuth(
		config.Auth,
		userRepo,
		tokenRepo,
	)

	collection := services.NewCollection(
		animeCollectionRepo,
		collectionRepo,
	)

	user := services.NewUser(
		userRepo,
	)

	// Create router.
	router := router.New(
		action,
		anime,
		auth,
		collection,
		user,
	)

	// Create server.
	srv := server.NewServer(config.Server, router)

	// Run server.
	go func() {
		err := srv.Run()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	err := srv.Shutdown()
	if err != nil {
		logger.Errorf("Error occurred on server shutting down", err)
	}

	db.Close()
}
