package domain_user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

var (
	UserNotFoundError     = errors.New("User not found")
	UserAreadyExistsError = errors.New("User already exists")
	UserInvalidIDError    = errors.New("User invalid id")
)

type IAccountRepository interface {
	Create(aggregates.User) error
	GetAll() ([]aggregates.User, error)
	GetById(uuid.UUID) (aggregates.User, error)
	GetByEmail(value_objects.Email) (aggregates.User, error)
	UpdateById(uuid.UUID, aggregates.User) error
	DeleteById(uuid.UUID) error
}
