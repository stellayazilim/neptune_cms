package Auth

import (
	"time"

	"github.com/google/uuid"
)

type AccessTokenPayload struct {
	Audience   string    `json:"aud"`
	Issuer     string    `json:"issuer"`
	Subject    string    `json:"sub"`
	IssuedAt   time.Time `json:"iat"`
	Roles      []byte    `json:"roles"`
	Expiration time.Time `json:"exp"`
	NotBefore  time.Time `json:"nbf"`
	Jti        uuid.UUID `json:"jti"`
}

type RefreshTokenPayload struct {
	Audience   string    `json:"aud"`
	Issuer     string    `json:"issuer"`
	Subject    string    `json:"sub"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
	NotBefore  time.Time `json:"nbf"`
	Jti        uuid.UUID `json:"jti"`
}
