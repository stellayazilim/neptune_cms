package Middlewares

import (
	"encoding/json"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/o1egl/paseto"
	"github.com/stellayazilim/neptune.domain/Auth"
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
)

func PasetoMiddlewareFactory(configService Providers.ConfigService) fiber.Handler {
	return pasetoware.New(pasetoware.Config{
		SymmetricKey: []byte(configService.AccessTokenSymmetricKey),
		TokenPrefix:  configService.TokenPrefix,
		ContextKey:   configService.TokenContextKey,
		Validate: func(decrypted []byte) (interface{}, error) {
			payload := new(Auth.AccessTokenPayload)
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
			return c.Next()
		},
	})
}
