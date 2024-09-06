package config

import (
	"github.com/ilfey/hikilist-go/internal/config/database"
	"github.com/ilfey/hikilist-go/internal/config/hasher"
	"github.com/ilfey/hikilist-go/internal/config/logger"
	"github.com/ilfey/hikilist-go/internal/config/server"
	"github.com/ilfey/hikilist-go/internal/config/tokenizer"
	"time"
)

// AppConfig is the application configuration.
type AppConfig struct {
	Tokenizer *tokenizer.Config
	Hasher    *hasher.Config
	Database  *database.Config
	Server    *server.Config
	Logger    *logger.Config
}

// New is the AppConfig constructor.
func New() *AppConfig {
	return &AppConfig{
		Tokenizer: &tokenizer.Config{
			Salt:            []byte(getEnv("TOKENIZER_SALT", "secret")),
			Issuer:          getEnv("TOKENIZER_ISSUER", "hikilist"),
			AccessLifeTime:  getEnvAsInt("TOKENIZER_ACCESS_LIFE_TIME", 24),
			RefreshLifeTime: getEnvAsInt("TOKENIZER_REFRESH_LIFE_TIME", 7*24),
		},
		Hasher: &hasher.Config{
			Salt: []byte(getEnv("HASHER_SALT", "secret")),
		},
		Database: &database.Config{
			Driver:   getEnv("DATABASE_DRIVER", "sqlite"),
			User:     getEnv("DATABASE_USER", ""),
			Password: getEnv("DATABASE_PASSWORD", ""),
			Host:     getEnv("DATABASE_HOST", ""),
			Port:     getEnv("DATABASE_PORT", ""),
			Database: getEnv("DATABASE_DATABASE", ""),
		},
		Server: &server.Config{
			ReadTimeout:       time.Duration(getEnvAsInt("SERVER_READ_TIMEOUT", 10000)),
			WriteTimeout:      time.Duration(getEnvAsInt("SERVER_WRITE_TIMEOUT", 10000)),
			IdleTimeout:       time.Duration(getEnvAsInt("SERVER_IDLE_TIMEOUT", 30000)),
			ReadHeaderTimeout: time.Duration(getEnvAsInt("SERVER_READ_HEADER_TIMEOUT", 2000)),

			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnv("SERVER_PORT", "5000"),
		},
		Logger: &logger.Config{
			ErrorBufferCap:   getEnvAsInt("LOGGER_ERROR_BUFFER_CAP", 5),
			RequestBufferCap: getEnvAsInt("LOGGER_REQUEST_BUFFER_CAP", 5),
		},
	}
}

func (AppConfig) GetEnv() Environment {
	return environment
}
