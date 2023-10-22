package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	interface_rest_common "github.com/stellayazilim/neptune_cms/internal/interfaces/rest/common"
	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest/common/middleware"
)

type IUserHandler interface {
	GetAll(*fiber.Ctx) error
}

type userHandler struct {
	interface_rest_common.BaseHandler
}

func UserHandler(base interface_rest_common.BaseHandler) IUserHandler {
	h := new(userHandler)
	h.BaseHandler = base
	return h
}

func InitUserRouter(a *fiber.App) error {

	r := a.Group("/users")

	b, err := BaseHandlerFactory(
		AddAuthService(),
		AddUserService())
	if err != nil {
		// handle error
		return err
	}
	h := UserHandler(b)

	r.Get(
		"/",
		middleware.PasetoMiddlewareFactory(
			os.Getenv("PASETO_TOKEN_PREFIX"),
			os.Getenv("PASETO_ACCESS_SYMMETRIC_KEY")),
		h.GetAll)

	return nil
}

func (h *userHandler) GetAll(ctx *fiber.Ctx) error {

	response, err := h.Services.User.GetAll()

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return ctx.JSON(response.Body)
}
