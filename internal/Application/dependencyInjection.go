package Application

import (
	"github.com/stellayazilim/neptune_cms/internal/Application/Auth/RegisterCommand"
	"go.uber.org/dig"
)

func UseApplication(c *dig.Container) {

	RegisterCommand.RegisterRegisterHandler(c)
}
