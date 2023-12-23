package Events

import (
	UserCreatedEventHandlers "github.com/stellayazilim/neptune.application/User/Events/Created"
	"go.uber.org/dig"
)

func RegisterUserEvents(c *dig.Container) {
	UserCreatedEventHandlers.ProvideUserCreatedEventHandlers(c)
}
