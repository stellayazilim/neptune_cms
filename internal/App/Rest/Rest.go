package Rest

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/stellayazilim/neptune_cms/internal/App/Rest/Handlers"
	"go.uber.org/dig"
)

func UseRest(c *dig.Container) {
	// provide fiber app
	c.Provide(fiber.New)

	// provide handlers
	c.Provide(AuthHandler)
	c.Provide(UserHandler)

	// init routers
	c.Invoke(AuthRouter)
	c.Invoke(UserRouter)
}

func Bootstrap(addr string) func(app *fiber.App) {
	return func(app *fiber.App) {
		app.Listen(addr)
	}
}
