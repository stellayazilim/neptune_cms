package Rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	. "github.com/stellayazilim/neptune.app/Rest/Handlers"
	"go.uber.org/dig"
)

func UseRest(c *dig.Container) {
	// provide fiber app
	c.Provide(fiber.New)

	c.Invoke(func(app *fiber.App) {

		app.Use(cors.New(cors.Config{

			Next:             nil,
			AllowOriginsFunc: nil,
			AllowOrigins:     "http://localhost:4321",
			AllowMethods:     "*",
			AllowHeaders:     "content-type, accept",
			AllowCredentials: true,
			ExposeHeaders:    "*",
			MaxAge:           0,
		}))
	})
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
