package rest_converter

import domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"

// todo convert domain errors to interface errors
func AuthErrorConverter() {}

func LoginDtoConverter(dto struct {
	Email    string
	password []byte
}) domain_auth.LoginDto {
	return domain_auth.LoginDto{
		Email:    dto.Email,
		Password: dto.password,
	}
}
