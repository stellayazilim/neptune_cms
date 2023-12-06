package ValueObjects

import (
	"github.com/google/uuid"
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Common/Models"
)

type RoleID struct {
	ValueObject[uuid.UUID]
}

func NewRoleID() RoleID {
	return RoleID{
		ValueObject: NewValueObject[uuid.UUID](uuid.New()),
	}
}
