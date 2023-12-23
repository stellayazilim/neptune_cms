package UserCreatedEventHandlers

import (
	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune.domain/User"
	"go.uber.org/dig"
)

func ProvideUserCreatedEventHandlers(c *dig.Container) {
	c.Provide(ProvideVerificationCodeEventHandler)

	c.Invoke(func(handler1 *VerificationCodeEventHandler) {

		mediatr.RegisterNotificationHandler[*User.UserCreatedEvent](handler1)
	})
}
