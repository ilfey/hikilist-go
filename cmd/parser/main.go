package main

import (
	"context"
	"net/http"
	"time"

	"github.com/ilfey/hikilist-go/internal/httpx"
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
	"github.com/ilfey/hikilist-go/pkg/config"
	"github.com/ilfey/hikilist-go/pkg/database"
	"github.com/ilfey/hikilist-go/pkg/parser"
	"github.com/ilfey/hikilist-go/pkg/parser/shiki/service"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/ilfey/hikilist-go/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/ttys3/rotatefilehook"
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

	animeRepo := repositories.NewAnime(db)

	// Create services.

	animeService := services.NewAnime(
		logger.WithField("service", "anime"),

		animeRepo,
	)

	shiki := service.New(
		shikiBuilder(cfg),
	)

	parser := &parser.Parser{
		Shiki: shiki,
		Anime: animeService,
	}

	result := <-parser.Run(context.Background())

	if result.Error() != nil {
		logger.Fatal(result.Error())
	}
}

func shikiBuilder(cfg *config.Config) *requestbuilder.RequestBuilder {
	builder := requestbuilder.NewRequestBuilder(
		cfg.Shiki.BaseUrl+"/api/",
		&http.Client{
			Transport: http.DefaultTransport,
			Timeout:   2000 * time.Millisecond,
		},
	)

	builder.AddResponseHook(func(r *httpx.Response) {
		logrus.Trace(r)
	})

	builder.AddRequestHook(func(rb *httpx.RequestBuilder) {
		logrus.Tracef("New request: %s", rb)
	})

	return builder
}
