package services

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	domain_user_mem "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
)

// todo user service

type IUserService interface {
	GetAll()
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

func UserServiceWithMemUserRepository(s *UserService) error {
	users := make(map[uuid.UUID]aggregates.User)

	userRepository := domain_user_mem.New(users)
	s.Repositories.User = userRepository
	return nil
}

func (s *UserService) GetAll() {}
