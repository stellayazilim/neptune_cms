package Entities

import (
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Common/Models"
	. "github.com/stellayazilim/neptune_cms/internal/Domain/User/ValueObjects"
)

type UserEntity struct {
	Entity[UserID]

	firstName string
	lastName  string
	password  []byte
	email     string
}

func NewUser(
	firstName string,
	lastName string,
	password []byte,
	email string,
) UserEntity {
	return UserEntity{
		Entity: Entity[UserID]{
			ID: NewUserID(),
		},
		firstName: firstName,
		lastName:  lastName,
		password:  password,
		email:     email,
	}
}
