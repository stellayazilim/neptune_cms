package rest

import (
	"github.com/gofiber/fiber/v2"
)

type IRest interface {
	Run(string) error
}

type rest struct {
	router *fiber.App
}

func Rest() IRest {
	return &rest{
		router: fiber.New(),
	}
}

// start rest application and listen given addr
func (r *rest) Run(addr string) error {
	return r.router.Listen(addr)
}
