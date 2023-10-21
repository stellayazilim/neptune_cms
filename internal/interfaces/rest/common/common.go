package interface_rest_common

import "github.com/stellayazilim/neptune_cms/pkg/services"

type BaseHandlerFactoryCfg func(*BaseHandler) error
type BaseHandler struct {
	Services struct {
		Auth services.IAuthService
		User services.IUserService
	}
}
