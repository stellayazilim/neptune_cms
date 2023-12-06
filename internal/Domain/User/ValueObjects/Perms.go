package ValueObjects

import (
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Common/Models"
)

type Perms struct {
	ValueObject[[]byte]
}

func NewPerms(perms []byte) Perms {
	return Perms{
		ValueObject: NewValueObject[[]byte](perms),
	}
}
