package rest_converter

import (
	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest/dto"
	domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

// todo convert domain errors to interface errors
func AuthErrorConverter() {}

func LoginDtoConverter(dto dto.LoginDto) domain_auth.LoginDto {

	return domain_auth.LoginDto{
		Email:    dto.Email,
		Password: value_objects.Password(dto.Password),
	}
}

func RegisterDtoConverter(dto dto.RegisterDto) domain_auth.RegisterDto {
	return domain_auth.RegisterDto{
		Email:    dto.Email,
		Password: value_objects.Password(dto.Password),
	}
}
