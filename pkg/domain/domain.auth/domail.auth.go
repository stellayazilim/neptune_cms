package domain_auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

var PasswordInvalidErr = errors.New("password invalid")
var TokenInvalidErr = errors.New("token invalid")
var TokenExpiredErr = errors.New("token expired")

type LoginDto struct {
	Email    string
	Password []byte
}
type RegisterDto struct {
	Email    string
	Password string
}

type PasetoPayload struct {
	Audience   string    `json:"aud"`
	Issuer     string    `json:"issuer"`
	Subject    string    `json:"sub"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
	NotBefore  time.Time `json:"nbf"`
	Jti        uuid.UUID `json:"jti"`
}

func CreatePasetoPayload(
	email string,
	duration time.Duration) *PasetoPayload {

	now := time.Now()

	return &PasetoPayload{
		Audience:   "Account",
		Issuer:     "Stella",
		Subject:    email,
		IssuedAt:   now,
		Expiration: now.Add(duration),
		NotBefore:  now,
		Jti:        uuid.New(),
	}
}

func (p *PasetoPayload) CreatePasetoTokenByPayload(key []byte) (string, error) {

	return paseto.NewV2().Encrypt(key, p, nil)
}
