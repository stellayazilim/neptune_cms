package Gorm

import (
	"github.com/stellayazilim/neptune.infrastructure/Common/Providers"
	"github.com/stellayazilim/neptune.infrastructure/Persistence/Gorm/Models"
	"gorm.io/gorm"
)

func UseMigration(db *gorm.DB, config *Providers.ConfigService) {
	if config.IsDevolopment {
		db.AutoMigrate(&Models.UserModel{})
	}
}
