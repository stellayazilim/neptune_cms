package PasswordResetCommand

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune.application/Common/Interfaces"
	"github.com/stellayazilim/neptune.application/Repositories"
	"go.uber.org/dig"
)

type PasswordResetHandler struct {
	userRepository Repositories.IUserRepository
	hashProvider   Interfaces.IHashProvider
}

func (h *PasswordResetHandler) Handle(
	ctx context.Context,
	command *PasswordResetCommand,
) (*PasswordResetCommandResponse, error) {

	h.userRepository.UpdatePasswordByEmail(
		command.Session.Subject,
		h.hashProvider.GenHash(command.NewPassword))

	return &PasswordResetCommandResponse{}, nil
}

func RegisterPasswordResetHandler(c *dig.Container) {

	c.Provide(func(
		userRepository Repositories.IUserRepository,
		hashProvider Interfaces.IHashProvider) *PasswordResetHandler {
		return &PasswordResetHandler{
			userRepository: userRepository,
			hashProvider:   hashProvider,
		}
	})

	c.Invoke(func(handler *PasswordResetHandler) {
		mediatr.RegisterRequestHandler[*PasswordResetCommand, *PasswordResetCommandResponse](handler)
	})
}
