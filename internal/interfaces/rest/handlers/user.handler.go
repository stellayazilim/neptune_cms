package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune_cms/pkg/services"
)

type IUserHandler interface {
}

type UserHandler struct {
	Services struct {
		authService services.IAuthService
	}
}

func InitUserRouter(a *fiber.App) error {

	r := a.Group("/auth")

	b, err := BaseHandlerFactory(AddAuthService)
	if err != nil {
		// handle error
		return err
	}
	h := AuthHandler(b)

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	return nil
}
