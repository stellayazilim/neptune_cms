package bcrypt

import (
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hash, plain value_objects.Password) bool {
	if err := bcrypt.CompareHashAndPassword(hash, plain); err != nil {
		return false
	}
	return true
}

func GenHash(p value_objects.Password) (value_objects.Password, error) {
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	return hash, err

}
