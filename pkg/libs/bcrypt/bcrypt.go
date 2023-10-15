package bcrypt

import (
	account_entity "github.com/stellayazilim/neptune_cms/pkg/entities/account.entity"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hash, plain account_entity.Password) bool {
	if err := bcrypt.CompareHashAndPassword(hash, plain); err != nil {
		return false
	}
	return true
}

func GenHash(p *account_entity.Password) error {
	hash, err := bcrypt.GenerateFromPassword(*p, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	*p = hash

	return nil
}
