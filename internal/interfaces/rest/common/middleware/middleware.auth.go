package middleware

import (
	"encoding/json"
	"fmt"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/o1egl/paseto"
	domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

func PasetoMiddlewareFactory(tokenPrefix, symmetricKey string) fiber.Handler {

	return pasetoware.New(pasetoware.Config{
		SymmetricKey: []byte(symmetricKey),
		TokenPrefix:  tokenPrefix,

		Validate: func(decrypted []byte) (interface{}, error) {
			payload := new(domain_auth.PasetoPayload)
			if err := json.Unmarshal(decrypted, payload); err != nil {
				return payload, err
			}

			return payload, nil
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			switch err {
			case paseto.ErrDataUnmarshal:
				return fiber.ErrUnprocessableEntity
			case paseto.ErrUnsupportedTokenType:
				return fiber.ErrUnprocessableEntity
			case paseto.ErrUnsupportedTokenVersion:
				return fiber.ErrUnprocessableEntity
			default:
				return fiber.ErrUnauthorized
			}
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			fmt.Println(c.Locals(c.Locals("user")))
			return c.Next()
		},
	})
}

func PermGuard(perms value_objects.Perms) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Locals(pasetoware.DefaultContextKey))
		return nil
	}
}
