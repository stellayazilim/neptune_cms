package main

import (
	"github.com/stellayazilim/neptune_cms/internal/App/Rest"
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure"
	env "github.com/stellayazilim/neptune_cms/internal/Infrastructure/Common/Env"
	"go.uber.org/dig"
)

func main() {
	// init .env files
	env := env.EnvProvider()
	env.Provide("./env/.env")

	// create ioc container
	container := dig.New()

	// Use rest api
	Infrastructure.UseInfrastructure(container)
	Rest.UseRest(container)

	// Start rest api
	container.Invoke(Rest.Bootstrap(":8080"))
}
