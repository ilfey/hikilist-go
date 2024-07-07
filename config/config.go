package config

import (
	authConfig "github.com/ilfey/hikilist-go/config/auth"
	databaseConfig "github.com/ilfey/hikilist-go/config/database"
	"github.com/ilfey/hikilist-go/internal/server"
)

// Глобальный конфиг приложения
type Config struct {
	Auth     *authConfig.Config
	Database *databaseConfig.Config
	Server   *server.Config
}

// Конструктор глобального конфига приложения
func NewConfig() *Config {
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
			DBName:   getEnv("DB_CONFIG_DBNAME", "hiki.db"),
		},
		Server: &server.Config{
			Host: getEnv("SERVER_CONFIG_HOST", "0.0.0.0"),
			Port: getEnvAsInt("SERVER_CONFIG_PORT", 5000),
		},
	}
}
