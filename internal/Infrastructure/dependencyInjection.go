package Infrastructure

import (
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Common/Providers"
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Persistence/Gorm"
	"go.uber.org/dig"
)

func UseInfrastructure(c *dig.Container) {

	c.Provide(Providers.ConfigServiceProvider)

	c.Provide(Gorm.GormProvider)
	Gorm.UseGormRepositoryProvider(c)
	c.Invoke(Gorm.UseMigration)
}
