package LogoutCommand

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune.application/Repositories"
	"github.com/stellayazilim/neptune.domain/Auth"
	"go.uber.org/dig"
)

type LogoutCommandDeps struct {
	dig.In
	TokenRepository Repositories.ITokenRepository
}

type LogoutCommandHandler struct {
	tokenRepository Repositories.ITokenRepository
}

func (h *LogoutCommandHandler) Handle(
	ctx context.Context,
	command *LogoutCommand) (
	*LogoutCommandResponse, error) {

	h.tokenRepository.BlacklistAccessToken(ctx, command.Token, ctx.Value("session").(Auth.AccessTokenPayload))
	return &LogoutCommandResponse{}, nil
}

func RegisterLogoutHandler(c *dig.Container) {

	c.Provide(func(
		h LogoutCommandDeps,
	) *LogoutCommandHandler {

		return &LogoutCommandHandler{
			tokenRepository: h.TokenRepository,
		}
	})

	c.Invoke(func(handler *LogoutCommandHandler) {
		mediatr.RegisterRequestHandler[
			*LogoutCommand,
			*LogoutCommandResponse](handler)
	})
}
