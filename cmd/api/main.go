package main

import (
	"github.com/ilfey/hikilist-go/internal/app"
	config2 "github.com/ilfey/hikilist-go/internal/config"
	"github.com/ilfey/hikilist-go/internal/infrastucture/di/container"
	"sync"
)

func main() {

	config2.MustLoadEnvironment()

	cfg := config2.New()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go app.NewApp(
		cfg,
		container.NewServiceContainerManager(),
	).Run(wg)

	wg.Wait()

	//
	//app.NewApp(cfg).Run()
	//
	//logger := createLogger(env)
	//
	//db := database.New(cfg.Database)
	//
	//// Create repositories.
	//
	//actionRepo := repositories.NewAction(db)
	//animeRepo := repositories.NewAnime(db)
	//animeCollectionRepo := repositories.NewAnimeCollection(db)
	//collectionRepo := repositories.NewCollection(db, actionRepo)
	//userRepo := repositories.NewUser(db, actionRepo)
	//tokenRepo := repositories.NewToken(db)
	//
	//// Create service.
	//
	//hasher := security.NewBcryptService(cfg.Tokenizer)
	//
	//jwtService := tokenizer.NewJwtService(
	//	logger.WithField("service", "tokenizer"),
	//	cfg.Tokenizer,
	//	tokenRepo,
	//)
	//
	//actionService := action.NewAction(
	//	logger.WithField("service", "action"),
	//
	//	actionRepo,
	//)
	//
	//animeService := anime.NewAnime(
	//	logger.WithField("service", "anime"),
	//
	//	animeRepo,
	//)
	//
	//authService := auth.NewAuth(
	//	logger.WithField("service", "auth"),
	//
	//	cfg.Tokenizer,
	//	hasher,
	//	jwtService,
	//	userRepo,
	//)
	//
	//collectionService := collection.NewCollection(
	//	logger.WithField("service", "collection"),
	//
	//	animeCollectionRepo,
	//	collectionRepo,
	//)
	//
	//userService := user.NewUser(
	//	logger.WithField("service", "user"),
	//
	//	userRepo,
	//)
	//
	//// Create router.
	//r := router.New(
	//	logger,
	//
	//	actionService,
	//	animeService,
	//	authService,
	//	collectionService,
	//	userService,
	//	jwtService,
	//)
	//
	//// Create server.
	//srv := server.NewServer(cfg.Server, r)
	//
	//// Listen server.
	//go func() {
	//	err := srv.Listen()
	//	if err != nil {
	//		logger.Fatal(err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//
	//<-quit
	//
	//err := srv.Shutdown()
	//if err != nil {
	//	logger.Errorf("Error occurred on server shutting down %v", err)
	//}
	//
	//db.Close()
}
