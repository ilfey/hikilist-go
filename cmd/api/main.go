package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ilfey/hikilist-go/internal/server"
	"github.com/ilfey/hikilist-go/pkg/api/router"
	"github.com/ilfey/hikilist-go/pkg/config"
	"github.com/ilfey/hikilist-go/pkg/database"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/sirupsen/logrus"
	"github.com/ttys3/rotatefilehook"

	"github.com/ilfey/hikilist-go/pkg/services"
)

func loadConfig() (config.Environment, *config.Config) {
	env := config.MustLoadEnvironment()

	cfg := config.New()

	return env, cfg
}

func createLogger(env config.Environment) *logrus.Logger {
	logger := logrus.New()

	var lvl logrus.Level
	var formatter logrus.Formatter

	switch env {
	case config.EnvironmentDev:
		lvl = logrus.TraceLevel

		formatter = &logrus.TextFormatter{
			ForceColors: true,
			ForceQuote:  true,
		}

	case config.EnvironmentProd:
		lvl = logrus.InfoLevel

		formatter = &logrus.JSONFormatter{
			PrettyPrint: true,
		}
	}

	logger.SetFormatter(formatter)
	logger.SetLevel(lvl)

	logrus.SetFormatter(formatter)
	logrus.SetLevel(lvl)

	// Add file hook in production.
	if env.IsProduction() {
		rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
			Filename:   "logs/api.log",
			MaxSize:    50, // Mb
			MaxBackups: 3,
			MaxAge:     7, // Days
			Level:      lvl,
			Formatter:  formatter,
		})
		if err != nil {
			logger.Fatal(err)
		}

		logger.AddHook(rotateFileHook)
		logrus.AddHook(rotateFileHook)
	}

	return logger
}

func main() {
	env, cfg := loadConfig()

	logger := createLogger(env)

	db := database.New(cfg.Database)

	// Create repositories.

	actionRepo := repositories.NewAction(db)
	animeRepo := repositories.NewAnime(db)
	animeCollectionRepo := repositories.NewAnimeCollection(db)
	collectionRepo := repositories.NewCollection(db, actionRepo)
	userRepo := repositories.NewUser(db, actionRepo)
	tokenRepo := repositories.NewToken(db)

	// Create services.

	action := services.NewAction(
		logger.WithField("service", "action"),

		actionRepo,
	)

	anime := services.NewAnime(
		logger.WithField("service", "anime"),

		animeRepo,
	)

	auth := services.NewAuth(
		logger.WithField("service", "auth"),

		cfg.Auth,
		userRepo,
		tokenRepo,
	)

	collection := services.NewCollection(
		logger.WithField("service", "collection"),

		animeCollectionRepo,
		collectionRepo,
	)

	user := services.NewUser(
		logger.WithField("service", "user"),

		userRepo,
	)

	// Create router.
	router := router.New(
		logger,

		action,
		anime,
		auth,
		collection,
		user,
	)

	// Create server.
	srv := server.NewServer(cfg.Server, router)

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
		logger.Errorf("Error occurred on server shutting down %v", err)
	}

	db.Close()
}
