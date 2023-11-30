package ValueObjects

import (
	"github.com/google/uuid"
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Common/Models"
)

type UserID struct {
	ValueObject[uuid.UUID]
}

func NewUserID() UserID {
	return UserID{
		ValueObject: NewValueObject[uuid.UUID](uuid.New()),
	}
}
