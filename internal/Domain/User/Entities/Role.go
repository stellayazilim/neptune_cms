package Entities

import (
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Common/Models"
	. "github.com/stellayazilim/neptune_cms/internal/Domain/User/ValueObjects"
)

type RoleEntity struct {
	Entity[RoleID]

	Name  string
	Perms Perms
}
