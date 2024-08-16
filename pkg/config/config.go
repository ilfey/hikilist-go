package config

import (
	"time"

	"github.com/ilfey/hikilist-go/internal/server"
	"github.com/ilfey/hikilist-go/pkg/config/auth"
	"github.com/ilfey/hikilist-go/pkg/config/database"
	"github.com/ilfey/hikilist-go/pkg/parser/shikimori/shiki"
)

// Глобальный конфиг приложения
type Config struct {
	Auth      *auth.Config
	Database  *database.Config
	Server    *server.Config
	Shikimori *shiki.Config
}

// Конструктор глобального конфига приложения
func New() *Config {
	return &Config{
		Auth: &auth.Config{
			Secret:          []byte(getEnv("AUTH_CONFIG_SECRET", "")),
			Issuer:          getEnv("AUTH_CONFIG_ISSUER", "hikilist"),
			AccessLifeTime:  getEnvAsInt("AUTH_CONFIG_ACCESS_LIFE_TIME", 24),
			RefreshLifeTime: getEnvAsInt("AUTH_CONFIG_REFRESH_LIFE_TIME", 7*24),
		},
		Database: &database.Config{
			Driver:   getEnv("DB_CONFIG_DRIVER", "sqlite"),
			User:     getEnv("DB_CONFIG_USER", ""),
			Password: getEnv("DB_CONFIG_PASSWORD", ""),
			Host:     getEnv("DB_CONFIG_HOST", ""),
			Port:     getEnvAsInt("DB_CONFIG_PORT", 0),
			Database: getEnv("DB_CONFIG_DATABASE", "hiki.db"),
		},
		Server: &server.Config{
			ReadTimeout:       time.Duration(getEnvAsInt("SERVER_CONFIG_READ_TIMEOUT", 10000)),
			WriteTimeout:      time.Duration(getEnvAsInt("SERVER_CONFIG_WRITE_TIMEOUT", 10000)),
			IdleTimeout:       time.Duration(getEnvAsInt("SERVER_CONFIG_IDLE_TIMEOUT", 30_000)),
			ReadHeaderTimeout: time.Duration(getEnvAsInt("SERVER_CONFIG_READ_HEADER_TIMEOUT", 2000)),

			Host: getEnv("SERVER_CONFIG_HOST", "0.0.0.0"),
			Port: getEnvAsInt("SERVER_CONFIG_PORT", 5000),
		},
		Shikimori: &shiki.Config{
			BaseUrl: getEnv("PARSER_SHIKIMORI_BASE_URL", "https://shikimori.one"),
		},
	}
}
