package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	interface_rest_common "github.com/stellayazilim/neptune_cms/internal/interfaces/rest/common"
	"github.com/stellayazilim/neptune_cms/pkg/common/dto"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
)

type IAuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
}

type authHandler struct {
	interface_rest_common.BaseHandler
}
type AuthHandlerFactoryCf func(*authHandler) error

func AuthHandler(base interface_rest_common.BaseHandler) IAuthHandler {
	h := new(authHandler)
	h.BaseHandler = base
	return h
}

func InitAuthRouter(a *fiber.App) error {

	r := a.Group("/auth")

	b, err := BaseHandlerFactory(AddAuthService())

	if err != nil {
		// todo handle err

		return err
	}
	h := AuthHandler(b)

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	return nil
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {

	body := new(dto.LoginRequest)

	if err := ctx.BodyParser(body); err != nil {
		// todo parse error for status code

		return err
	}

	tokens, err := h.Services.Auth.Login(*body)

	if err != nil {
		// todo parse error for status code
		if errors.Is(err, domain_user.UserNotFoundError) {
			return fiber.ErrUnauthorized
		}

		return fiber.ErrUnauthorized
	}
	fmt.Println("login 3")
	return ctx.JSON(tokens)

}
func (h *authHandler) Register(ctx *fiber.Ctx) error {

	body := new(dto.RegisterRequest)
	if err := ctx.BodyParser(body); err != nil {
		// todo parse error for status code
		return fiber.ErrUnprocessableEntity
	}

	if err := h.Services.Auth.Register(*body); err != nil {
		return fiber.ErrConflict
	}

	return nil
}
