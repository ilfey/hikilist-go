package databaseConfig

// Конфиг базы данных
type Config struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}
