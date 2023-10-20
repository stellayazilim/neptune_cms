package handlers

import (
	"errors"

	"github.com/stellayazilim/neptune_cms/pkg/services"
)

var ErrAuthServiceAlreadyExist = errors.New("auth service already exist")

type BaseHandlerFactoryCfg func(*baseHandler) error
type baseHandler struct {
	Services struct {
		Auth services.IAuthService
		User services.IUserService
	}
}

func BaseHandlerFactory(configs ...BaseHandlerFactoryCfg) (baseHandler, error) {

	base := new(baseHandler)
	for _, cfg := range configs {
		if err := cfg(base); err != nil {
			return *base, err
		}
	}

	return *base, nil
}

func AddAuthService(h *baseHandler) error {
	s, err := services.AuthServiceFactory(services.AuthServiceWithMemUserRepository)
	if err != nil {
		return err
	}
	h.Services.Auth = s
	return nil
}
