package handlers

import "github.com/stellayazilim/neptune_cms/pkg/neptune/auth"

type handlerServices struct {
	authService auth.IAuthService
}

func HandlerServices(
	authService auth.IAuthService,
) *handlerServices {
	return &handlerServices{
		authService: authService,
	}
}
