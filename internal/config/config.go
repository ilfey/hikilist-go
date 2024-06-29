package config

import "<package_name>/internal/config/database"

// Global config
type Config struct {
	Database *database.Config
}

func NewConfig() *Config {
	return &Config{
		Database: &database.Config{
			User:     getEnv("DB_CONFIG_USER", "root"),
			Password: getEnv("DB_CONFIG_PASSWORD", "root"),
			Host:     getEnv("DB_CONFIG_HOST", "localhost"),
			Port:     getEnvAsInt("DB_CONFIG_PORT", 3306),
			DbName:   getEnv("DB_CONFIG_DBNAME", ""),
		},
	}
}