package database

import "fmt"

// Config represents database configuration.
type Config struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// DSN returns database connection string.
func (config *Config) DSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Driver,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
