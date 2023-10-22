package handlers

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"

	interface_rest_common "github.com/stellayazilim/neptune_cms/internal/interfaces/rest/common"
	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest/common/middleware"
	domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
)

type IAuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
	Refresh(*fiber.Ctx) error
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
	r.Post(
		"/refresh",
		middleware.PasetoMiddlewareFactory(
			os.Getenv("PASETO_TOKEN_PREFIX"),
			os.Getenv("PASETO_REFRESH_SYMMETRIC_KEY")),
		h.Refresh,
	)
	return nil
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {

	body := new(domain_auth.LoginRequest)

	if err := ctx.BodyParser(&body.Body); err != nil {
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
	return ctx.JSON(tokens)

}
func (h *authHandler) Register(ctx *fiber.Ctx) error {

	req := new(domain_auth.RegisterRequest)
	if err := ctx.BodyParser(&req.Body); err != nil {
		// todo parse error for status code
		return fiber.ErrUnprocessableEntity
	}

	if err := h.Services.Auth.Register(*req); err != nil {
		return fiber.ErrConflict
	}

	return nil
}

func (h *authHandler) Refresh(ctx *fiber.Ctx) error {
	session := interface_rest_common.GetSession(ctx)
	request := new(domain_auth.LoginRequest)
	request.Body.Email = session.Subject
	tokens, err := h.Services.Auth.Refresh(*request)

	if err != nil {
		// todo parse error for status code
		if errors.Is(err, domain_user.UserNotFoundError) {
			return fiber.ErrUnauthorized
		}

		return fiber.ErrUnauthorized
	}
	return ctx.JSON(tokens)
}
