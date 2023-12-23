package Setup

import (
	"github.com/stellayazilim/neptune.application/Common/Interfaces"
	"github.com/stellayazilim/neptune.application/Repositories"
	"github.com/stellayazilim/neptune.domain/User"
	UserEntities "github.com/stellayazilim/neptune.domain/User/Entities"
	"github.com/stellayazilim/neptune.domain/User/ValueObjects"
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
	"go.uber.org/dig"
)

type SetupDeps struct {
	dig.In
	ConfigService  *Providers.ConfigService
	HashProvider   Interfaces.IHashProvider
	UserRepository Repositories.IUserRepository
}

func Setup(deps SetupDeps) {

	role := UserEntities.NewRole(
		"SuperAdmin",
		ValueObjects.NewPerms([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16}))

	deps.UserRepository.Create(
		User.CreateUserAggregate(
			deps.ConfigService.CmsAdminFirstName,
			deps.ConfigService.CmsAdminLastName,
			deps.ConfigService.CmsAdminEmail,
			deps.HashProvider.GenHash([]byte(deps.ConfigService.CmsAdminPassword)),
			[]*UserEntities.RoleEntity{
				&role,
			},
		),
	)
}
