package config

import "github.com/joho/godotenv"

func InjectEnv() error {

	return godotenv.Load()
}
