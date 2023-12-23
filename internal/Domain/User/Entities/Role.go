package UserEntities

import (
	"github.com/stellayazilim/neptune.domain/Common/Models"
	"github.com/stellayazilim/neptune.domain/User/ValueObjects"
)

type RoleEntity struct {
	Models.Entity[ValueObjects.RoleID]

	Name  string
	Perms ValueObjects.Perms
}

func NewRole(name string, perms ValueObjects.Perms) RoleEntity {
	return RoleEntity{
		Entity: Models.Entity[ValueObjects.RoleID]{
			ID: ValueObjects.NewRoleID(),
		},
		Name:  name,
		Perms: perms,
	}

}
