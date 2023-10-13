package auth

import (
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type IAuthHelper interface {
	ComparePassword(acc *models.Account, dto *SigninDto) bool
	HashPassword(dto *SignupDto) []byte
}
type authHelper struct{}

func AuthHelper() IAuthHelper {
	return &authHelper{}
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
