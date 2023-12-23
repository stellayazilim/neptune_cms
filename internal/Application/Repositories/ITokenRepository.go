package Repositories

import (
	"context"

	"github.com/stellayazilim/neptune.domain/Auth"
)

type ITokenRepository interface {
	BlacklistAccessToken(context context.Context, token string, payload Auth.AccessTokenPayload)
	BlacklistRefreshToken(context context.Context, token string, payload Auth.RefreshTokenPayload)
	IsBlacklisted(ctx context.Context, token string) bool
}
