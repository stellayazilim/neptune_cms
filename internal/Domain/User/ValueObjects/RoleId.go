package ValueObjects

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune.domain/Common/Models"
)

type RoleID struct {
	Models.ValueObject[uuid.UUID]
}

func NewRoleID() RoleID {
	return RoleID{
		ValueObject: Models.NewValueObject[uuid.UUID](uuid.New()),
	}
}
