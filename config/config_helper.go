package config

import (
	"os"
	"strconv"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load("./configs/local.env")
	if err != nil {
		logger.Fatal(err)
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
