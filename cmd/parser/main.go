package main

import (
	"fmt"
	"os"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/pkg/config"
	"github.com/ilfey/hikilist-go/pkg/database"
	"github.com/ilfey/hikilist-go/pkg/parser"
	shikiService "github.com/ilfey/hikilist-go/pkg/parser/shikimori"
)

func main() {
	logger.SetLevel(logger.LevelDebug)

	config.LoadEnvironment()

	config := config.New()

	database.New(config.Database)

	shikiService := shikiService.NewShikimoriService(
		config.Shikimori,
	)

	parser := &parser.Parser{
		Shiki: shikiService,
	}

	err := parser.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
