package entities

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

type Role struct {
	ID    uuid.NullUUID
	Name  string
	Perms value_objects.Perms
}
