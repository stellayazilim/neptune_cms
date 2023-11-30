package env

import (
	"github.com/joho/godotenv"
	. "github.com/stellayazilim/neptune_cms/internal/Application/Common/Interfaces"
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

func EnvProvider() IEnvProvider {
	return &envProvider{}
}
