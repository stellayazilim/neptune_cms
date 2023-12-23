package Repositories

import "github.com/stellayazilim/neptune.domain/User"

type IUserRepository interface {
	Create(User.UserAggregate) error
	FindByEmail(email string) (User.UserAggregate, error)
	UpdatePasswordByEmail(string, []byte) error
}
