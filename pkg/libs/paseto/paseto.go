package paseto

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	account_entity "github.com/stellayazilim/neptune_cms/pkg/entities/account.entity"
)

var PasswordInvalid = errors.New("password_invalid")
var TokenInvalid = errors.New("token_invalid")
var TokenExpired = errors.New("token_expired")

type PasetoPayload struct {
	Audience   string                `json:"aud"`
	Issuer     string                `json:"issuer"`
	Subject    *account_entity.Email `json:"sub"`
	IssuedAt   time.Time             `json:"iat"`
	Expiration time.Time             `json:"exp"`
	NotBefore  time.Time             `json:"nbf"`
	Jti        uuid.UUID             `json:"jti"`
}

func CreatePasetoPayload(
	email *account_entity.Email,
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
