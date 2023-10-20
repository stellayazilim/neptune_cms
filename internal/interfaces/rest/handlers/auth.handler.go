package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	rest_converter "github.com/stellayazilim/neptune_cms/internal/interfaces/rest/converters"
	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest/dto"
	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest/serializers"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	"github.com/stellayazilim/neptune_cms/pkg/services"
)

type IAuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
}

type authHandler struct {
	baseHandler
	services struct {
		authService services.IAuthService
	}
}
type AuthHandlerFactoryCf func(*authHandler) error

func AuthHandler(base baseHandler) IAuthHandler {
	h := new(authHandler)
	h.baseHandler = base
	return h
}

func InitAuthRouter(a *fiber.App) error {

	r := a.Group("/auth")

	b, err := BaseHandlerFactory(AddAuthService)

	if err != nil {
		// todo handle err
	}
	h := AuthHandler(b)

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	return nil
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {

	body := new(dto.LoginDto)

	if err := ctx.BodyParser(body); err != nil {
		// todo parse error for status code

		return err
	}

	tokens, err := h.services.authService.Login(rest_converter.LoginDtoConverter(*body))

	if err != nil {
		// todo parse error for status code

		fmt.Println(errors.Is(err, domain_user.UserNotFoundError))
		if errors.Is(err, domain_user.UserNotFoundError) {
			return fiber.ErrUnauthorized
		}

		return fiber.ErrUnauthorized
	}
	fmt.Println("login 3")
	return ctx.JSON(serializers.LoginResponseSerializer{
		AccessToken:  tokens[0],
		RefreshToken: tokens[1],
	})

}
func (h *authHandler) Register(ctx *fiber.Ctx) error {

	body := new(dto.RegisterDto)
	if err := ctx.BodyParser(body); err != nil {
		// todo parse error for status code
		return fiber.ErrUnprocessableEntity
	}

	if err := h.services.authService.Register(rest_converter.RegisterDtoConverter(*body)); err != nil {
		return fiber.ErrConflict
	}

	return nil
}
