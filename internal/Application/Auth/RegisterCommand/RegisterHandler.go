package RegisterCommand

import (
	"context"
	"fmt"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/stellayazilim/neptune_cms/internal/Application/Repositories"
	"go.uber.org/dig"
)

type RegisterHandler struct {
	userRepository Repositories.IUserRepository
}

func (h *RegisterHandler) Handle(
	ctx context.Context,
	command RegisterCommand) (RegisterCommandResponse, error) {

	fmt.Println("register")
	return RegisterCommandResponse{}, nil
}

func RegisterRegisterHandler(c *dig.Container) {

	c.Provide(func(userRepository Repositories.IUserRepository) *RegisterHandler {
		return &RegisterHandler{
			userRepository: userRepository,
		}
	})

	c.Invoke(func(handler *RegisterHandler) {
		mediatr.RegisterRequestHandler[RegisterCommand, RegisterCommandResponse](handler)
	})

}
