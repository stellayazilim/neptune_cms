package auth

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type IAuthHelper interface {
	ComparePassword(acc *models.Account, dto *SigninDto) bool
	HashPassword(dto *SignupDto) []byte
	CreateToken(key []byte, p *models.Account, duration time.Duration) string
	Paseto() *paseto.V2
}
type authHelper struct {
	paseto paseto.V2
}

func AuthHelper() IAuthHelper {
	return &authHelper{
		paseto: *paseto.NewV2(),
	}
}

func (h *authHelper) ComparePassword(acc *models.Account, dto *SigninDto) bool {
	if err := bcrypt.CompareHashAndPassword(acc.Password, []byte(dto.Password)); err != nil {
		return false
	}
	return true
}

func (h *authHelper) HashPassword(dto *SignupDto) []byte {

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		// handle hash errors
	}
	return hash
}

func (h *authHelper) CreateToken(key []byte, p *models.Account, duration time.Duration) string {
	now := time.Now()
	payload := PasetoPayload{
		Audience:   "Account",
		Issuer:     "api.neptunecms.com",
		Subject:    p.Email,
		IssuedAt:   now,
		Expiration: now.Add(duration),
		NotBefore:  now,
		Jti:        uuid.New().String(),
	}

	token, err := h.paseto.Encrypt(key, payload, nil)

	if err != nil {
		log.Fatal(err)
	}
	return token

}

func (h *authHelper) Paseto() *paseto.V2 {
	return &h.paseto
}
