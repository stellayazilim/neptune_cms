package handlers

import (
	"errors"

	"github.com/stellayazilim/neptune_cms/pkg/services"
)

var AuthServiceAlreadyExist = errors.New("Auth service already exist")

type BaseHandlerFactoryCfg func(*baseHandler) error
type baseHandler struct {
	services struct {
		auth services.IAuthService
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
	h.services.auth = s
	return nil
}
