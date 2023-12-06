package Gorm

import (
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Common/Providers"
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Persistence/Gorm/Models"
	"gorm.io/gorm"
)

func UseMigration(db *gorm.DB, config *Providers.ConfigService) {
	if config.IsDevolopment {
		db.AutoMigrate(&Models.UserModel{})
	}
}
