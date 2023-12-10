package Infrastructure

import (
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
	"github.com/stellayazilim/neptune.infrastructure/Persistence/Gorm"
	"go.uber.org/dig"
)

func UseInfrastructure(c *dig.Container) {

	c.Provide(Providers.HashProvider)
	c.Provide(Providers.ConfigServiceProvider)
	c.Provide(Providers.EnvProvider)
	c.Provide(Providers.PasetoTokenProvider)
	c.Provide(Gorm.GormProvider)
	c.Provide(Providers.DateTimeProvider)

	Gorm.UseGormRepositoryProvider(c)
	c.Invoke(Gorm.UseMigration)
}
