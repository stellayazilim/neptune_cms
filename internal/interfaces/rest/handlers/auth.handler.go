package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest/dto"
	"github.com/stellayazilim/neptune_cms/pkg/services"
)

type IAuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
}

type authHandler struct {
	services struct {
		authService services.IAuthService
	}
}

func AuthHandler(services ...func(*authHandler) error) IAuthHandler {
	h := new(authHandler)
	for _, service := range services {
		if err := service(h); err != nil {
			log.Fatal(err)
		}
	}
	return h
}

func AddAuthService(h *authHandler) error {
	if h.services.authService != nil {
		return AuthServiceAlreadyExist
	}
	return nil
}

func InitAuthRouter(a *fiber.App) error {

	r := a.Group("/auth")
	h := AuthHandler(
		AddAuthService,
	)

	r.Post("/", h.Register)
	r.Post("/", h.Login)
	return nil
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {

	dto := new(dto.LoginDto)
	if err := ctx.BodyParser(dto); err != nil {
		// todo parse error for status code
		return err
	}

	return nil
}
func (h *authHandler) Register(ctx *fiber.Ctx) error {
	return nil
}
