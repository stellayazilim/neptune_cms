package ValueObjects

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune.domain/Common/Models"
)

type UserID struct {
	Models.ValueObject[uuid.UUID]
}

func NewUserID() UserID {
	return UserID{
		ValueObject: Models.NewValueObject[uuid.UUID](uuid.New()),
	}
}

func ParseUserID(id uuid.UUID) UserID {
	return UserID{
		ValueObject: Models.NewValueObject[uuid.UUID](id),
	}
}
