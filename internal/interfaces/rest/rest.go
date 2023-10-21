package rest

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest/handlers"
)

// factory token for init handler on rest interface
type RestInitHandler = func(*fiber.App) error

// factory token add service to a handler
type HandlerServiceCfg[H any] func(H) error

type IRest interface {
	Run(string) error
	Stop() error
}

type rest struct {
	App *fiber.App
}

func RestFactory(handlers ...RestInitHandler) IRest {

	r := &rest{
		App: fiber.New(),
	}

	for _, h := range handlers {
		if err := h(r.App); err != nil {
			log.Fatal(err)
		}
	}
	// init  handlers
	return r
}

func NewRestWithHandlers() IRest {
	return RestFactory(
		handlers.InitAuthRouter,
		handlers.InitUserRouter,
	)
}

// start rest application and listen given addr
func (r *rest) Run(addr string) error {
	return r.App.Listen(addr)
}

func (r *rest) Stop() error {
	return r.App.Shutdown()
}
