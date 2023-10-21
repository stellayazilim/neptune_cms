package domain_user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAreadyExists = errors.New("user already exists")
	ErrUserInvalidID    = errors.New("user invalid id")
)

type IUserRepository interface {
	Create(aggregates.User) error
	GetAll() ([]aggregates.User, error)
	GetById(uuid.UUID) (aggregates.User, error)
	GetByEmail(value_objects.Email) (aggregates.User, error)
	UpdateById(uuid.UUID, aggregates.User) error
	DeleteById(uuid.UUID) error
}
