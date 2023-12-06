package Gorm

import (
	"fmt"

	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Common/Providers"
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Persistence/Gorm/Repositories"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormProvider(configService *Providers.ConfigService) *gorm.DB {
	fmt.Println("err")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Istanbul",
		configService.PostgresHost,
		configService.PostgresUser,
		configService.PostgresPassword,
		configService.PostgresDatabase,
		configService.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("err")
	return db

}

func UseGormRepositoryProvider(c *dig.Container) {

	c.Provide(Repositories.GormUserRepository)
}
