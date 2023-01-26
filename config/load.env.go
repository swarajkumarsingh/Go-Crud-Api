package config

import (
	"github.com/joho/godotenv"
	errorhandler "github.com/swarajkumarsingh/go-build/errorHandler"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	errorhandler.HandleError(err)
}
