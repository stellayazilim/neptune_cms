package UserCreatedEventHandlers

import (
	"context"

	"github.com/stellayazilim/neptune.domain/User"
	"go.uber.org/dig"
)

type (
	VerificationCodeEventHandlerDeps struct {
		dig.In
	}
	VerificationCodeEventHandler struct {
	}
)

func (h *VerificationCodeEventHandler) Handle(ctx context.Context, event *User.UserCreatedEvent) error {
	return nil
}

func ProvideVerificationCodeEventHandler(opts VerificationCodeEventHandlerDeps) *VerificationCodeEventHandler {
	return &VerificationCodeEventHandler{}
}
