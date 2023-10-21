package services

import (
	"github.com/stellayazilim/neptune_cms/pkg/common/dto"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	domain_user_mem "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
)

// todo user service

type IUserService interface {
	GetAll() (dto.UsersResponse, error)
}

type UserService struct {
	Repositories struct {
		User domain_user.IUserRepository
	}
}

func UserServiceFactory(cfgs ...ServiceConfig[UserService]) (IUserService, error) {

	as := new(UserService)
	for _, cfg := range cfgs {
		if err := cfg(as); err != nil {
			return as, err
		}
	}
	return as, nil
}

func UserServiceWithMemUserRepository() ServiceConfig[UserService] {
	return func(s *UserService) error {
		userRepository := domain_user_mem.New()
		s.Repositories.User = userRepository
		return nil
	}
}

func (s *UserService) GetAll() (dto.UsersResponse, error) {
	users, err := s.Repositories.User.GetAll()

	if err != nil {
		return *new(dto.UsersResponse), err
	}

	data := make([]dto.UserResponseBody, 0)

	for _, us := range users {
		data = append(data, dto.UserResponseBody{
			Email: string(us.GetAccount().Email),
		})
	}

	response := dto.UsersResponse{
		Body: dto.UsersResponseBody{
			Data:         data,
			Count:        1,
			Displaying:   1,
			TotalPage:    1,
			CurrentPage:  1,
			PreviousPage: "",
			NextPage:     "",
		},
	}

	return response, nil

}
