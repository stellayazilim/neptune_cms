package MeCommand

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune.application/Repositories"
	"go.uber.org/dig"
)

type MeCommandHandler struct {
	userRepository Repositories.IUserRepository
}

func (h *MeCommandHandler) Handle(
	ctx context.Context,
	command *MeCommand) (*MeCommandResponse, error) {

	user, err := h.userRepository.FindByEmail(command.Email)

	if err != nil {
		return &MeCommandResponse{}, err
	}
	return &MeCommandResponse{
		User: user,
	}, nil
}

func RegisterMeCommandHandler(c *dig.Container) {
	c.Provide(func(
		userRepository Repositories.IUserRepository,
	) *MeCommandHandler {
		return &MeCommandHandler{
			userRepository: userRepository,
		}
	})

	c.Invoke(func(handler *MeCommandHandler) {
		mediatr.RegisterRequestHandler[*MeCommand, *MeCommandResponse](handler)
	})
}
