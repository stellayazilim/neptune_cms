package interface_rest_common

import (
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"
	"github.com/stellayazilim/neptune_cms/pkg/services"
)

type BaseHandlerFactoryCfg func(*BaseHandler) error
type BaseHandler struct {
	Services struct {
		Auth services.IAuthService
		User services.IUserService
	}
}

func GetSession(ctx *fiber.Ctx) *domain_auth.PasetoPayload {
	return ctx.Locals(pasetoware.DefaultContextKey).(*domain_auth.PasetoPayload)
}
