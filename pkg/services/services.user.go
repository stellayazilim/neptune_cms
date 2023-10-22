package services

import (
	"fmt"

	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	domain_user_mem "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
)

// todo user service

type IUserService interface {
	GetAll() (domain_user.UsersResponse, error)
	GetById()
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

func (s *UserService) GetAll() (domain_user.UsersResponse, error) {
	users, err := s.Repositories.User.GetAll()

	if err != nil {
		return *new(domain_user.UsersResponse), err
	}

	data := make([]*domain_user.UsersResponseData, 0)

	for _, us := range users.Data {
		fmt.Println("user", us.GetAccount().Email)
		data = append(data, &domain_user.UsersResponseData{
			Email: string(us.GetAccount().Email),
			ID:    us.GetAccount().ID.UUID.String(),
		})
	}

	response := domain_user.UsersResponse{
		Body: domain_user.UsersResponseBody{
			Data:    &data,
			Current: uint64(len(data)),
			Total:   users.Total,
		},
	}

	fmt.Println(response.Body.Data)
	return response, nil

}

func (s *UserService) GetById() {}
