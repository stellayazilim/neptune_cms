package User

import (
	"errors"

	"github.com/stellayazilim/neptune.domain/Common/Models"
	UserEntities "github.com/stellayazilim/neptune.domain/User/Entities"
)

var ErrUserNotFoundException = errors.New("user not found")
var ErrUserAlreadyExistException = errors.New("user already exists")

type (
	UserAggregate struct {
		Models.AggregateRoot[UserEntities.UserEntity]
		Roles []*UserEntities.RoleEntity
	}

	UserCreatedEvent struct {
		User UserAggregate
	}

	UserPasswordUpdateEvent struct {
		User UserAggregate
	}

	UserValidationEvent struct {
		User UserAggregate
	}

	UserEmailUpdateEvent struct {
		User UserAggregate
	}
)

func CreateUserAggregate(
	FirstName string,
	LastName string,
	Email string,
	Password []byte,
	Roles []*UserEntities.RoleEntity,
) UserAggregate {
	userAggregate := new(UserAggregate)
	userAggregate.AggregateRoot.Root = UserEntities.NewUser(
		FirstName,
		LastName,
		Password,
		Email,
	)
	userAggregate.Roles = Roles
	return *userAggregate
}

func EmptyUserAggregate() UserAggregate {
	return UserAggregate{
		AggregateRoot: Models.AggregateRoot[UserEntities.UserEntity]{
			Root: *new(UserEntities.UserEntity),
		},
	}
}
