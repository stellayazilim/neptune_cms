package Application

import (
	"github.com/stellayazilim/neptune.application/Auth/Commands/LoginCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/LogoutCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/MeCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/PasswordResetCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/RegisterCommand"
	"github.com/stellayazilim/neptune.application/User/Events"
	"go.uber.org/dig"
)

func UseApplication(c *dig.Container) {
	handlers := []func(*dig.Container){
		// register commands
		RegisterCommand.RegisterRegisterHandler,
		LoginCommand.RegisterLoginHandler,
		PasswordResetCommand.RegisterPasswordResetHandler,
		LogoutCommand.RegisterLogoutHandler,
		MeCommand.RegisterMeCommandHandler,
		// register events
		Events.RegisterUserEvents,
	}

	for _, handler := range handlers {
		handler(c)
	}

}
