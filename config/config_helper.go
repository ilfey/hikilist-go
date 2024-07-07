package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func LoadEnvironment() {
	err := godotenv.Load("configs/local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
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