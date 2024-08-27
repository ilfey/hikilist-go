package config

import (
	"time"

	"github.com/ilfey/hikilist-go/internal/server"
	"github.com/ilfey/hikilist-go/pkg/config/auth"
	"github.com/ilfey/hikilist-go/pkg/config/database"
)

// Глобальный конфиг приложения
type Config struct {
	Auth     *auth.Config
	Database *database.Config
	Server   *server.Config
}

// Конструктор глобального конфига приложения
func New() *Config {
	return &Config{
		Auth: &auth.Config{
			Secret:          []byte(getEnv("AUTH_SECRET", "secret")),
			Issuer:          getEnv("AUTH_ISSUER", "hikilist"),
			AccessLifeTime:  getEnvAsInt("AUTH_ACCESS_LIFE_TIME", 24),
			RefreshLifeTime: getEnvAsInt("AUTH_REFRESH_LIFE_TIME", 7*24),
		},
		Database: &database.Config{
			Driver:   getEnv("DB_DRIVER", "sqlite"),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnvAsInt("DB_PORT", 0),
			Database: getEnv("DB_DATABASE", "hiki.db"),
		},
		Server: &server.Config{
			ReadTimeout:       time.Duration(getEnvAsInt("SERVER_READ_TIMEOUT", 10000)),
			WriteTimeout:      time.Duration(getEnvAsInt("SERVER_WRITE_TIMEOUT", 10000)),
			IdleTimeout:       time.Duration(getEnvAsInt("SERVER_IDLE_TIMEOUT", 30000)),
			ReadHeaderTimeout: time.Duration(getEnvAsInt("SERVER_READ_HEADER_TIMEOUT", 2000)),

			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvAsInt("SERVER_PORT", 5000),
		},
	}
}
