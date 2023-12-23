package LoginCommand

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune.domain/User"
)

type LoginCommand struct {
	Email    string `validate:"required,email"`
	Password []byte `validate:"required"`
}

type LoginCommandResponse struct {
	AccessToken  string
	RefreshToken string
	User         User.UserAggregate
}

type LoginCommandUserField struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  []byte
}
