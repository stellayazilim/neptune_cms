package Infrastructure

import (
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
	"github.com/stellayazilim/neptune.infrastructure/Persistence/Gorm"
	"github.com/stellayazilim/neptune.infrastructure/Persistence/Redis"
	"go.uber.org/dig"
)

func UseInfrastructure(c *dig.Container) {
	c.Provide(Gorm.GormProvider)
	c.Provide(Redis.UseRedisProvider)

	c.Provide(Providers.HashProvider)
	c.Provide(Providers.ConfigServiceProvider)
	c.Provide(Providers.EnvProvider)
	c.Provide(Providers.PasetoTokenProvider)

	c.Provide(Providers.DateTimeProvider)

	Gorm.UseGormRepositoryProvider(c)
	Redis.UseRedisRepositoryProviders(c)

	c.Invoke(Gorm.UseMigration)
}
