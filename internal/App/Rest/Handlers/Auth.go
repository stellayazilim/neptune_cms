package Handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune_cms/internal/Application/Auth/RegisterCommand"
	AuthContract "github.com/stellayazilim/neptune_cms/internal/Contracts/Auth"
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Common/Providers"
)

type authHandler struct{}

func AuthRouter(app *fiber.App, handler *authHandler, configService *Providers.ConfigService) {
	r := app.Group("/auth")
	r.Post("/login", handler.Login)
	r.Post("/register", handler.Register)
}

func AuthHandler() *authHandler {

	return &authHandler{}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {

	return ctx.SendString("Hello world")
}

func (h *authHandler) Register(ctx *fiber.Ctx) error {

	request := &AuthContract.RegisterRequestBody{}

	if err := ctx.JSON(request); err != nil {
		return err
	}
	fmt.Println("register controller")

	mediatr.Send[*RegisterCommand.RegisterCommand, *RegisterCommand.RegisterCommandResponse](ctx.Context(), &RegisterCommand.RegisterCommand{})
	return nil
}
