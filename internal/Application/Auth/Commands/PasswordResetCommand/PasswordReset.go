package PasswordResetCommand

import "github.com/stellayazilim/neptune.domain/Auth"

type PasswordResetCommand struct {
	NewPassword []byte `validate:"required"`
	Session     Auth.AccessTokenPayload
}

type PasswordResetCommandResponse struct {
}
