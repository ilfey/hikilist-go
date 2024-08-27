package main

import (
	"context"

	"github.com/ilfey/hikilist-go/pkg/config"
	"github.com/ilfey/hikilist-go/pkg/parser"
	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/source/shiki"
	shikiConfig "github.com/ilfey/hikilist-go/pkg/parser/source/shiki/config"
	"github.com/sirupsen/logrus"
	"github.com/ttys3/rotatefilehook"
)

// func loadConfig() (config.Environment, *config.Config) {
// 	env := config.MustLoadEnvironment()

// 	cfg := config.New()

// 	return env, cfg
// }

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
	// env, cfg := loadConfig()

	logger := createLogger(config.EnvironmentDev)

	shikiExtractor := shiki.New(
		logger,
		&shikiConfig.Config{
			BaseUrl:        "https://shikimori.one/",
			UserAgent:      "Hikilist",
			TickTimeout:    1000 * 60, // 1 minute
			RequestTimeout: 1000,      // 1 second
			Anime: &shikiConfig.AnimeConfig{
				Order:    "id_desc",
				Censored: "false",
				Limit:    50,
			},
		})

	parser := &parser.Parser{
		Logger: logger,
		Extractors: []extractor.Extractor{
			shikiExtractor,
		},
	}

	err := parser.Run(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
}
