package paseto

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/stellayazilim/neptune_cms/pkg/models"
)

var PasswordInvalid = errors.New("password_invalid")
var TokenInvalid = errors.New("token_invalid")
var TokenExpired = errors.New("token_expired")

type PasetoPayload struct {
	Audience   string    `json:"aud"`
	Issuer     string    `json:"issuer"`
	Subject    string    `json:"sub"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
	NotBefore  time.Time `json:"nbf"`
	Jti        uuid.UUID `json:"jti"`
}

func CreatePasetoPayload(p *models.Account, duration time.Duration) *PasetoPayload {

	now := time.Now()

	return &PasetoPayload{
		Audience:   "Account",
		Issuer:     "Stella",
		Subject:    p.Email,
		IssuedAt:   now,
		Expiration: now.Add(duration),
		NotBefore:  now,
		Jti:        uuid.New(),
	}
}

func (p *PasetoPayload) CreatePasetoTokemByPayload(key []byte) (string, error) {

	return paseto.NewV2().Encrypt(key, p, nil)
}
