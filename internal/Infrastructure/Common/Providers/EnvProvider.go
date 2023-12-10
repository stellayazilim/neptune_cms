package Providers

import (
	"github.com/joho/godotenv"
	"github.com/stellayazilim/neptune.application/Common/Interfaces"
)

type envProvider struct {
}

// Provide implements Interfaces.IEnvProvider.
func (*envProvider) Provide(path string) {
	err := godotenv.Load(path)

	if err != nil {
		panic(err)
	}
}

func EnvProvider() Interfaces.IEnvProvider {
	return &envProvider{}
}
