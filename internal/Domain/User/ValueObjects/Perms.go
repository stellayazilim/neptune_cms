package ValueObjects

import (
	"github.com/stellayazilim/neptune.domain/Common/Models"
)

type Perms struct {
	Models.ValueObject[[]byte]
}

func NewPerms(perms []byte) Perms {
	return Perms{
		ValueObject: Models.NewValueObject[[]byte](perms),
	}
}
