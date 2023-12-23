package LoginCommand

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune.application/Auth"
	"github.com/stellayazilim/neptune.application/Common/Interfaces"
	"github.com/stellayazilim/neptune.application/Repositories"
	DomainAuth "github.com/stellayazilim/neptune.domain/Auth"
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
	"go.uber.org/dig"
)

type LoginCommandHandlerDeps struct {
	dig.In
	UserRepository   Repositories.IUserRepository
	HashProvider     Interfaces.IHashProvider
	TokenProvider    Interfaces.ITokenProvider
	ConfigService    *Providers.ConfigService
	DateTimeProvider Interfaces.IDateTimeProvider
}

type LoginHandler struct {
	userRepository   Repositories.IUserRepository
	hashProvider     Interfaces.IHashProvider
	tokenProvider    Interfaces.ITokenProvider
	configService    *Providers.ConfigService
	dateTimeProvider Interfaces.IDateTimeProvider
}

func (h *LoginHandler) Handle(
	ctx context.Context,
	command *LoginCommand) (*LoginCommandResponse, error) {

	// check if user exist
	user, err := h.userRepository.FindByEmail(command.Email)
	if err != nil {
		return &LoginCommandResponse{}, err
	}

	// check password
	if !h.hashProvider.Compare(command.Password, user.Root.Password) {
		return &LoginCommandResponse{}, Auth.ErrPasswordDoesNotMatch
	}
	// create access token
	accessToken := h.tokenProvider.GenerateAccessToken(DomainAuth.AccessTokenPayload{
		Audience:  h.configService.TokenAudience,
		Issuer:    h.configService.TokenIssuer,
		Subject:   user.Root.Email,
		IssuedAt:  h.dateTimeProvider.UTCNow(),
		NotBefore: h.dateTimeProvider.UTCNow(),
		Expiration: h.dateTimeProvider.UTCNow().Add(
			time.Minute * h.configService.TokenAccessExpiration,
		),
		Jti: uuid.New(),
	})
	refreshToken := h.tokenProvider.GenerateRefreshToken(DomainAuth.RefreshTokenPayload{
		Audience:  h.configService.TokenAudience,
		Issuer:    h.configService.TokenIssuer,
		Subject:   user.Root.Email,
		IssuedAt:  h.dateTimeProvider.UTCNow(),
		NotBefore: h.dateTimeProvider.UTCNow(),
		Expiration: h.dateTimeProvider.UTCNow().Add(
			time.Minute * h.configService.TokenAccessExpiration,
		),
		Jti: uuid.New(),
	})
	// create refresh token
	return &LoginCommandResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil

}

func RegisterLoginHandler(c *dig.Container) {
	c.Provide(func(
		h LoginCommandHandlerDeps,
	) *LoginHandler {
		return &LoginHandler{
			userRepository:   h.UserRepository,
			hashProvider:     h.HashProvider,
			configService:    h.ConfigService,
			dateTimeProvider: h.DateTimeProvider,
			tokenProvider:    h.TokenProvider,
		}
	})

	c.Invoke(func(handler *LoginHandler) {
		mediatr.RegisterRequestHandler[*LoginCommand, *LoginCommandResponse](handler)
	})
}
