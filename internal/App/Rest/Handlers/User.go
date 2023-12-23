package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune.app/Rest/Middlewares"
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
)

type userHandler struct {
	ConfigService *Providers.ConfigService
}

func UserRouter(app *fiber.App, handler *userHandler) {
	r := app.Group("/users")
	r.Get("/", Middlewares.PasetoMiddlewareFactory(*handler.ConfigService), handler.GetAll)
}

func UserHandler(configService *Providers.ConfigService) *userHandler {

	return &userHandler{
		ConfigService: configService,
	}
}

func (h *userHandler) GetAll(ctx *fiber.Ctx) error {

	return ctx.JSON([]string{})
}
