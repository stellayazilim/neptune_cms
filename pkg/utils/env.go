package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func InjectEnv() error {

	switch os.Getenv("GO_ENV") {

	case "test":
		return godotenv.Load("../env/.env.test")
	default:
		return godotenv.Load()

	}

}
