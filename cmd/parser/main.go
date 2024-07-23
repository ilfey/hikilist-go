package main

import (
	"fmt"
	"os"

	"github.com/ilfey/hikilist-go/config"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/parser"
	shikiService "github.com/ilfey/hikilist-go/parser/shikimori"
	animeService "github.com/ilfey/hikilist-go/services/anime"
)

func main() {
	logger.SetLevel(logger.LevelDebug)

	config.LoadEnvironment()

	config := config.NewConfig()

	db := database.NewDatabase(config.Database)

	animeService := animeService.New(
		db,
	)

	shikiService := shikiService.NewShikimoriService(
		config.Shikimori,
	)

	parser := &parser.Parser{
		Anime: animeService,
		Shiki: shikiService,
	}

	err := parser.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
