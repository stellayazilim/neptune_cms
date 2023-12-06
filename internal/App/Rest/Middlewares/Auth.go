package Auth

import (
	"encoding/json"
	"fmt"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/o1egl/paseto"
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Auth"
	. "github.com/stellayazilim/neptune_cms/internal/Infrastructure/Common/Providers"
)

func PasetoMiddlewareFactory(configService ConfigService) fiber.Handler {
	return pasetoware.New(pasetoware.Config{
		SymmetricKey: []byte(configService.AccessTokenSymmetricKey),
		TokenPrefix:  configService.TokenPrefix,

		Validate: func(decrypted []byte) (interface{}, error) {
			payload := new(TokenPayload)
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
