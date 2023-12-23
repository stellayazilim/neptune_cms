package Interfaces

import "github.com/stellayazilim/neptune.domain/Auth"

type ITokenProvider interface {
	GenerateAccessToken(Auth.AccessTokenPayload) string
	GenerateRefreshToken(Auth.RefreshTokenPayload) string
}
