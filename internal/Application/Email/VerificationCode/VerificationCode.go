package VerificationCode

import (
	"context"

	"github.com/stellayazilim/neptune.domain/User"
	"go.uber.org/dig"
)

type VerificationCodeEventHandler struct {
}

func (e *VerificationCodeEventHandler) Handle(ctx context.Context, event *User.UserCreatedEvent) error {
	return nil
}

func RegisterVerificationCodeEmailEventHandler(c *dig.Container) {
	c.Provide(func() *VerificationCodeEventHandler {
		return &VerificationCodeEventHandler{}
	})
}
