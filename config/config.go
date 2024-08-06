package config

import (
	"time"

	authConfig "github.com/ilfey/hikilist-go/config/auth"
	databaseConfig "github.com/ilfey/hikilist-go/config/database"
	"github.com/ilfey/hikilist-go/internal/server"
	shikiConfig "github.com/ilfey/hikilist-go/parser/shikimori/config"
)

// Глобальный конфиг приложения
type Config struct {
	Auth      *authConfig.Config
	Database  *databaseConfig.Config
	Server    *server.Config
	Shikimori *shikiConfig.Config
}

// Конструктор глобального конфига приложения
func New() *Config {
	return &Config{
		Auth: &authConfig.Config{
			Secret: []byte(getEnv("AUTH_CONFIG_SECRET", "")),
			Issuer: getEnv("AUTH_CONFIG_ISSUER", "hikilist"),
		},
		Database: &databaseConfig.Config{
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
		Shikimori: &shikiConfig.Config{
			BaseUrl: getEnv("PARSER_SHIKIMORI_BASE_URL", "https://shikimori.one"),
		},
	}
}
