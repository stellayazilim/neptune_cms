package auth

import (
	"errors"
	"time"
)

var PasswordDoesNotMatchError = errors.New("Password does not match")

type PasetoPayload struct {
	Audience   string    `json:"aud"`
	Issuer     string    `json:"issuer"`
	Subject    string    `json:"sub"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
	NotBefore  time.Time `json:"nbf"`
	Jti        string    `json:"jti"`
}

type SigninTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
