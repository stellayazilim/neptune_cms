package main

import (
	"github.com/stellayazilim/neptune.app/Rest"
	Application "github.com/stellayazilim/neptune.application"
	"github.com/stellayazilim/neptune.application/Setup"
	Infrastructure "github.com/stellayazilim/neptune.infrastructure"
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
	"go.uber.org/dig"
)

func main() {
	// init .env files
	env := Providers.EnvProvider()
	env.Provide("./env/.env")

	// create ioc container
	container := dig.New()

	// Use rest api
	Infrastructure.UseInfrastructure(container)
	Application.UseApplication(container)

	Rest.UseRest(container)

	container.Invoke(Setup.Setup)
	// Start rest api
	container.Invoke(Rest.Bootstrap(":8080"))
}
