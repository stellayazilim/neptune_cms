package RegisterCommand

import (
	"context"
	"errors"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune.application/Common/Interfaces"
	"github.com/stellayazilim/neptune.application/Repositories"
	"github.com/stellayazilim/neptune.domain/User"
	UserEntities "github.com/stellayazilim/neptune.domain/User/Entities"
	"go.uber.org/dig"
)

type RegisterHandler struct {
	userRepository Repositories.IUserRepository
	hashProvider   Interfaces.IHashProvider
}

func (h *RegisterHandler) Handle(
	ctx context.Context,
	command *RegisterCommand) (*RegisterCommandResponse, error) {

	user := User.CreateUserAggregate(
		command.FirstName,
		command.LastName,
		command.Email,
		h.hashProvider.GenHash(command.Password),
		make([]*UserEntities.RoleEntity, 0),
	)

	if err := h.userRepository.Create(user); err != nil {
		return &RegisterCommandResponse{}, err
	}

	if err := mediatr.Publish[*User.UserCreatedEvent](ctx, &User.UserCreatedEvent{
		User: user,
	}); err != nil {
		return &RegisterCommandResponse{}, errors.New("verification email could not be sent")
	}

	return &RegisterCommandResponse{
		ID: user.Root.ID.GetValue(),
	}, nil
}

func RegisterRegisterHandler(c *dig.Container) {

	c.Provide(func(
		userRepository Repositories.IUserRepository,
		hashProvider Interfaces.IHashProvider) *RegisterHandler {
		return &RegisterHandler{
			userRepository: userRepository,
			hashProvider:   hashProvider,
		}
	})

	c.Invoke(func(handler *RegisterHandler) {
		mediatr.RegisterRequestHandler[*RegisterCommand, *RegisterCommandResponse](handler)
	})
}
