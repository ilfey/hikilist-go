package main

import (
	"context"
	"net/http"
	"time"

	"github.com/ilfey/hikilist-go/internal/httpx"
	"github.com/ilfey/hikilist-go/internal/logger"
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
	"github.com/ilfey/hikilist-go/pkg/config"
	"github.com/ilfey/hikilist-go/pkg/database"
	"github.com/ilfey/hikilist-go/pkg/parser"
	"github.com/ilfey/hikilist-go/pkg/parser/shiki/service"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/ilfey/hikilist-go/pkg/services"
)

func main() {
	logger.SetLevel(logger.LevelDebug)

	config.LoadEnvironment()

	config := config.New()

	db := database.New(config.Database)

	animeRepo := repositories.NewAnime(db)
	animeService := services.NewAnime(animeRepo)

	shiki := service.New(
		shikiBuilder(config),
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

func shikiBuilder(config *config.Config) *requestbuilder.RequestBuilder {
	builder := requestbuilder.NewRequestBuilder(
		config.Shiki.BaseUrl+"/api/",
		&http.Client{
			Transport: http.DefaultTransport,
			Timeout:   2000 * time.Millisecond,
		},
	)

	builder.AddResponseHook(func(r *httpx.Response) {
		logger.Trace(r)
	})

	builder.AddRequestHook(func(rb *httpx.RequestBuilder) {
		logger.Tracef("New request: %s", rb)
	})

	return builder
}
