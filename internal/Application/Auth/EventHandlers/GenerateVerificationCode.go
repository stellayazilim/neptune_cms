package AuthEventHandlers

import (
	"context"

	"github.com/stellayazilim/neptune.domain/User"
)

type (
	VerificationCodeEventHandler struct{}
)

func (h *VerificationCodeEventHandler) Handler(ctx context.Context, event User.UserCreatedEvent) error {

	return nil
}
