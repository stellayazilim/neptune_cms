package interface_rest_common

import (
	"github.com/stellayazilim/neptune_cms/pkg/services"
)

func AddUserServiceWithMemRepository() BaseHandlerFactoryCfg {

	return func(h *BaseHandler) error {
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
