package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune_cms/internal/rest/handlers"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
)

type IRest interface {
	Run(string) error
	Stop() error
}

type rest struct {
	App *fiber.App
}

func Rest(
	authService auth.IAuthService,
) IRest {

	r := &rest{
		App: fiber.New(),
	}

	s := handlers.HandlerServices(authService)

	handlers.AuthHandler(r.App.Group("/auth"), s)

	fmt.Println(r.App.GetRoutes())
	return r

}

// start rest application and listen given addr
func (r *rest) Run(addr string) error {
	return r.App.Listen(addr)
}

func (r *rest) Stop() error {
	return r.App.Shutdown()
}
