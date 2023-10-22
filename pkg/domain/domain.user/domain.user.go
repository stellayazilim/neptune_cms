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

type IUserRepository interface {
	Create(aggregates.User) error
	GetAll() (struct {
		Data  []aggregates.User
		Total uint64
	}, error)
	GetById(uuid.UUID) (aggregates.User, error)
	GetByEmail(value_objects.Email) (aggregates.User, error)
	UpdateById(uuid.UUID, aggregates.User) error
	DeleteById(uuid.UUID) error
}

type UsersResponseData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
type UsersResponseBody struct {
	Data        *[]*UsersResponseData `json:"data"`
	Total       uint64                `json:"total"`
	Current     uint64                `json:"displayCurrent"`
	CurrentPage uint8                 `json:"currentPage"`
	TotalPage   uint8                 `json:"totalPage"`
}

type UsersResponse struct {
	Body UsersResponseBody
}

type UserRequestByIdParams struct {
	Id string `param:"id"`
}

type UsersRequestParams struct{}
type UsersRequest struct {
}
