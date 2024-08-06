package databaseConfig

import "fmt"

// Конфиг базы данных
type Config struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func (config *Config) DSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		config.Driver,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
