package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Environment uint

const (
	EnvironmentDev Environment = iota
	EnvironmentProd
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
	err := godotenv.Load("./configs/development.env")
	if err != nil {
		return err
	}

	environment = EnvironmentDev

	return nil
}

func loadProd() error {
	err := godotenv.Load("./configs/production.env")
	if err != nil {
		return err
	}

	environment = EnvironmentProd

	return nil
}

func MustLoadEnvironment() Environment {
	err := loadProd()
	if err != nil {
		logrus.Infof("Error occurred while loading production.env %v", err)
	} else {
		return environment
	}

	err = loadDev()
	if err != nil {
		logrus.Fatal(err)
	}

	return environment
}
