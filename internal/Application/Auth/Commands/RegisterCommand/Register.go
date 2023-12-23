package RegisterCommand

import "github.com/google/uuid"

type RegisterCommand struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required"`
	Password  []byte `validate:"required"`
}

type RegisterCommandResponse struct {
	ID uuid.UUID
}
