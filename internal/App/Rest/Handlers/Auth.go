package Handlers

import (
	"github.com/gofiber/fiber/v2"
)

type authHandler struct{}

func AuthRouter(app *fiber.App, handler *authHandler) {
	r := app.Group("/auth")
	r.Post("/login", handler.Login)
	r.Post("/register", handler.Login)
}

func AuthHandler() *authHandler {

	return &authHandler{}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {

	return ctx.SendString("Hello world")
}
