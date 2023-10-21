package handlers

import (
	"errors"

	interface_rest_common "github.com/stellayazilim/neptune_cms/internal/interfaces/rest/common"
	"github.com/stellayazilim/neptune_cms/pkg/services"
)

var ErrAuthServiceAlreadyExist = errors.New("auth service already exist")

func BaseHandlerFactory(configs ...interface_rest_common.BaseHandlerFactoryCfg) (interface_rest_common.BaseHandler, error) {

	base := new(interface_rest_common.BaseHandler)
	for _, cfg := range configs {
		if err := cfg(base); err != nil {
			return *base, err
		}
	}

	return *base, nil
}
func AddAuthService() interface_rest_common.BaseHandlerFactoryCfg {
	return func(h *interface_rest_common.BaseHandler) error {
		s, err := services.AuthServiceFactory(
			services.AuthServiceWithMemUserRepository(),
		)
		if err != nil {
			return err
		}
		h.Services.Auth = s
		return nil
	}
}

func AddUserService() interface_rest_common.BaseHandlerFactoryCfg {

	return func(h *interface_rest_common.BaseHandler) error {
		s, err := services.UserServiceFactory(
			services.UserServiceWithMemUserRepository(),
		)

		if err != nil {
			return err
		}

		h.Services.User = s

		return nil
	}

}
