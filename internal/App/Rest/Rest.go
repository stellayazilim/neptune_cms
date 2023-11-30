package Rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	. "github.com/stellayazilim/neptune_cms/internal/App/Rest/Handlers"
	"go.uber.org/dig"
)

func UseRest(c *dig.Container) {
	// provide fiber app
	c.Provide(fiber.New)

	// provide handlers
	c.Provide(AuthHandler)

	fmt.Println("init routera")
	// init routers
	c.Invoke(AuthRouter)
}

func Bootstrap(addr string) func(app *fiber.App) {
	return func(app *fiber.App) {
		app.Listen(addr)
	}
}
