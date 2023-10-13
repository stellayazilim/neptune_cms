package utils

import (
	"github.com/joho/godotenv"
)

func InjectEnv(path ...string) error {

	return godotenv.Load(path...)
}
