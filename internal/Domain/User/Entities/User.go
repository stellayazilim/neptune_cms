package UserEntities

import (
	"github.com/stellayazilim/neptune.domain/Common/Models"
	"github.com/stellayazilim/neptune.domain/User/ValueObjects"
)

type UserEntity struct {
	Models.Entity[ValueObjects.UserID]

	FirstName string
	LastName  string
	Password  []byte
	Email     string
}

func NewUser(
	firstName string,
	lastName string,
	password []byte,
	email string,
) UserEntity {
	return UserEntity{
		Entity: Models.Entity[ValueObjects.UserID]{
			ID: ValueObjects.NewUserID(),
		},
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
		Email:     email,
	}
}

func (a *UserEntity) SetPassword(password []byte) {
	a.Password = password
}
