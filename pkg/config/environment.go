package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Environment string

const (
	EnvironmentDev  = "development"
	EnvironmentProd = "production"
)

func (env Environment) IsDevelopment() bool {
	return env == EnvironmentDev
}

func (env Environment) IsProduction() bool {
	return env == EnvironmentProd
}

var environment Environment

func GetEnv() Environment {
	return environment
}

func loadDev() error {
	err := godotenv.Load("./configs/local.env")
	if err != nil {
		return err
	}

	environment = EnvironmentDev

	return nil
}

func MustLoadEnvironment() Environment {
	err := loadDev()
	if err != nil {
		logrus.Fatal(err)
	}

	return environment
}
