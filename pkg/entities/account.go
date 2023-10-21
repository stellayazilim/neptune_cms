package entities

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

type Account struct {
	Base
	ID           uuid.NullUUID
	Email        value_objects.Email
	Password     value_objects.Password
	OldPasswords []value_objects.Password
}

func NewAccount() *Account {

	return &Account{
		Base:     Base{},
		ID:       uuid.NullUUID{},
		Email:    *new(value_objects.Email),
		Password: value_objects.Password{},
	}
}
