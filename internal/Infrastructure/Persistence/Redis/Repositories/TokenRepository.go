package Repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stellayazilim/neptune.application/Repositories"
	"github.com/stellayazilim/neptune.domain/Auth"
)

type tokenRepository struct {
	redis *redis.Client
}

func TokenRepository(redis *redis.Client) Repositories.ITokenRepository {
	return &tokenRepository{
		redis: redis,
	}
}

// Blacklist implements Repositories.ITokenRepository.
func (r *tokenRepository) BlacklistAccessToken(ctx context.Context, token string, payload Auth.AccessTokenPayload) {

	r.redis.SetEx(context.Background(), token, payload, time.Until(payload.Expiration))

}
func (r *tokenRepository) BlacklistRefreshToken(ctx context.Context, token string, payload Auth.RefreshTokenPayload) {

	r.redis.SetEx(context.Background(), token, payload, time.Until(payload.Expiration))
}

// IsBlacklisted implements Repositories.ITokenRepository.
func (r *tokenRepository) IsBlacklisted(ctx context.Context, token string) bool {
	if err := r.redis.Get(ctx, token); err != nil {
		return false
	}
	return true
}
