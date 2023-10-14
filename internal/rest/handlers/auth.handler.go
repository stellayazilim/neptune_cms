package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
)

type authHandler struct {
	services *handlerServices
}

func AuthHandler(router fiber.Router, services *handlerServices) {
	h := &authHandler{
		services: services,
	}

	router.Post("/signin", h.Signin)
	router.Post("/signup", h.Signup)
}

func (h *authHandler) Signin(ctx *fiber.Ctx) error {

	signDto := new(auth.SigninDto)
	if err := ctx.BodyParser(signDto); err != nil {
		return fiber.ErrUnprocessableEntity
	}
	tokens, err := h.services.authService.Signin(signDto)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrUnauthorized
	}
	return ctx.JSON(&auth.SigninTokenResponse{
		AccessToken:  tokens[0],
		RefreshToken: tokens[1],
	})
}

func (h *authHandler) Signup(ctx *fiber.Ctx) error {

	signupDto := new(auth.SignupDto)

	if err := ctx.BodyParser(signupDto); err != nil {
		return fiber.ErrUnprocessableEntity
	}

	if err := h.services.authService.Signup(signupDto); err != nil {
		return fiber.ErrConflict
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
