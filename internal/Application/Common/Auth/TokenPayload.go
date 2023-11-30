package Auth

import (
	"time"

	"github.com/google/uuid"
)

type TokenPayload struct {
	Audience   string    `json:"aud"`
	Issuer     string    `json:"issuer"`
	Subject    string    `json:"sub"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
	NotBefore  time.Time `json:"nbf"`
	Jti        uuid.UUID `json:"jti"`
}

func (p *TokenPayload) CreatePayload(
	audience string,
	issuer string,
	subject string,
	issuedAt time.Time,
	expiration time.Time,
	notBefore time.Time,
	jti uuid.UUID,
) TokenPayload {
	return TokenPayload{
		Audience:   audience,
		Issuer:     issuer,
		Subject:    subject,
		IssuedAt:   issuedAt,
		Expiration: expiration,
		NotBefore:  notBefore,
		Jti:        jti,
	}
}
