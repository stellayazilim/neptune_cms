package Handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Common/Providers"
)

type userHandler struct {
	configService *Providers.ConfigService
}

func UserRouter(app *fiber.App, handler *userHandler) {
	fmt.Println("User router")
	r := app.Group("/users")
	r.Get("/", handler.GetAll)
}

func UserHandler(configService *Providers.ConfigService) *userHandler {

	return &userHandler{
		configService: configService,
	}
}

func (h *userHandler) GetAll(ctx *fiber.Ctx) error {

	return ctx.JSON([]string{})
}
