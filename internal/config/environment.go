package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"io/fs"
)

type Environment uint64

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
		if !errors.Is(err, fs.ErrNotExist) {
			panic("Error occurred while loading production.env " + err.Error())
		}
	} else {
		return environment
	}

	err = loadDev()
	if err != nil {
		panic("neither development.env nor production.env found")
	}

	return environment
}
