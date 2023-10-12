package rest

import (
	"github.com/gofiber/fiber/v2"
)

type IRest interface {
	Run(string) error
	Stop() error
}

type rest struct {
	App *fiber.App
}

func Rest() IRest {
	return &rest{
		App: fiber.New(),
	}
}

// start rest application and listen given addr
func (r *rest) Run(addr string) error {
	return r.App.Listen(addr)
}

func (r *rest) Stop() error {
	return r.App.Shutdown()
}
