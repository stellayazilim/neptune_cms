package Handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune.app/Rest/Middlewares"
	"github.com/stellayazilim/neptune.application/Auth"
	"github.com/stellayazilim/neptune.application/Auth/Commands/LoginCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/LogoutCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/MeCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/PasswordResetCommand"
	"github.com/stellayazilim/neptune.application/Auth/Commands/RegisterCommand"
	LoginContract "github.com/stellayazilim/neptune.contracts/Auth/Login"
	MeContract "github.com/stellayazilim/neptune.contracts/Auth/Me"
	PasswordContract "github.com/stellayazilim/neptune.contracts/Auth/Password"
	RegisterContract "github.com/stellayazilim/neptune.contracts/Auth/Register"
	"github.com/stellayazilim/neptune.contracts/ProblemDetail"
	DomainAuth "github.com/stellayazilim/neptune.domain/Auth"
	"github.com/stellayazilim/neptune.domain/User"
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
)

type authHandler struct {
	configService *Providers.ConfigService
	validator     *validator.Validate
}

func AuthRouter(app *fiber.App, handler *authHandler, configService *Providers.ConfigService) {
	r := app.Group("/auth")
	r.Post("/login", handler.Login)
	r.Post("/register", handler.Register)
	r.Post("/logout", Middlewares.PasetoMiddlewareFactory(*configService), handler.Logout)
	r.Patch("/password", Middlewares.PasetoMiddlewareFactory(*configService), handler.ResetPassword)
	r.Get("/me", Middlewares.PasetoMiddlewareFactory(*configService), handler.Me)
}

func AuthHandler(configService *Providers.ConfigService) *authHandler {
	return &authHandler{
		configService: configService,
	}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {

	requestBody := &LoginContract.LoginRequestBody{}
	if err := ctx.BodyParser(requestBody); err != nil {
		ctx.Status(422)
		return ctx.JSON(fiber.Map{
			"errors": err.Error(),
		}, "application/problem+json")
	}
	response, err := mediatr.Send[
		*LoginCommand.LoginCommand,
		*LoginCommand.LoginCommandResponse](
		ctx.Context(),
		&LoginCommand.LoginCommand{
			Email:    requestBody.Email,
			Password: []byte(requestBody.Password),
		})

	if err != nil && errors.Is(err, Auth.ErrPasswordDoesNotMatch) {
		ctx.Status(fiber.ErrUnauthorized.Code)
		return ctx.JSON(ProblemDetail.New(
			"Invalid Credentials",
			string(ctx.Request().URI().Path()),
			401,
			"Authentication",
		), "application/problem+json")
	}

	perms := make([]string, 0)

	for _, role := range response.User.Roles {
		for _, perm := range role.Perms.GetValue() {
			perms = append(perms, strconv.Itoa(int(perm)))
		}

	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "neptune",
		Value:    response.AccessToken,
		HTTPOnly: true,
		Path:     "/",
		Secure:   true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "neptune_reserved",
		Value:    response.RefreshToken,
		HTTPOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return ctx.JSON(LoginContract.LoginResponseBody{
		Email:        response.User.Root.Email,
		ID:           response.User.Root.ID.GetValue().String(),
		FirstName:    response.User.Root.FirstName,
		LastName:     response.User.Root.LastName,
		Perms:        perms,
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})

}

func (h *authHandler) Register(ctx *fiber.Ctx) error {

	request := &RegisterContract.RegisterRequestBody{}
	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	if errs := request.Validate(ctx); errs != nil {
		ctx.Status(fiber.ErrUnprocessableEntity.Code)
		return ctx.JSON(errs, "application/problem+json; charset=utf8")
	}

	response, err := mediatr.Send[
		*RegisterCommand.RegisterCommand,
		*RegisterCommand.RegisterCommandResponse](
		ctx.Context(),
		&RegisterCommand.RegisterCommand{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Password:  []byte(request.Password),
		})

	if errors.Is(err, User.ErrUserAlreadyExistException) {
		ctx.Status(fiber.ErrConflict.Code)
		return ctx.JSON(
			ProblemDetail.New(
				string(ctx.Request().URI().Path()),
				"User already exist",
				409,
				"Duplicate entity",
			), "application/problem+json; charset=utf8")
	}
	return ctx.JSON(response)
}

func (h *authHandler) ResetPassword(ctx *fiber.Ctx) error {

	request := PasswordContract.PasswordResetRequestBody{}

	session := ctx.Locals(h.configService.TokenContextKey).(*DomainAuth.AccessTokenPayload)
	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(422)
		ctx.JSON(fiber.Error{
			Code:    422,
			Message: err.Error(),
		})
	}

	_, err := mediatr.Send[
		*PasswordResetCommand.PasswordResetCommand,
		*PasswordResetCommand.PasswordResetCommandResponse](
		ctx.Context(),
		&PasswordResetCommand.PasswordResetCommand{
			NewPassword: []byte(request.NewPassword),
			Session:     *session,
		})

	if err != nil {
		return err
	}

	return nil
}

func (h *authHandler) Logout(ctx *fiber.Ctx) error {

	token := strings.Replace(ctx.Get("Authorization"), "Bearer ", "", 1)


	mediatr.Send[*LogoutCommand.LogoutCommand, *LoginCommand.LoginCommandResponse](
		ctx.Context(),
		&LogoutCommand.LogoutCommand{
			Token: token,
		})
	return nil
}

func (h *authHandler) Me(ctx *fiber.Ctx) error {
	session := ctx.Locals(h.configService.TokenContextKey).(*DomainAuth.AccessTokenPayload)

	me, err := mediatr.Send[*MeCommand.MeCommand, *MeCommand.MeCommandResponse](
		ctx.Context(),
		&MeCommand.MeCommand{
			Email: session.Subject,
		},
	)

	if err != nil {
		return nil
	}

	perms := make([]string, 0)

	for _, role := range me.User.Roles {
		for _, perm := range role.Perms.GetValue() {
			perms = append(perms, strconv.Itoa(int(perm)))
		}

	}
	return ctx.JSON(&MeContract.MeResponseBody{
		Id:        me.User.Root.ID.GetValue().String(),
		FirstName: me.User.Root.FirstName,
		LastName:  me.User.Root.LastName,
		Email:     me.User.Root.Email,
		Perms:     perms,
	})
}
