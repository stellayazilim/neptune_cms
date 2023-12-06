package Repositories

import (
	"github.com/stellayazilim/neptune_cms/internal/Application/Repositories"
	"github.com/stellayazilim/neptune_cms/internal/Domain/User"
	"github.com/stellayazilim/neptune_cms/internal/Infrastructure/Persistence/Gorm/Models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func GormUserRepository(db *gorm.DB) Repositories.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user User.UserAggregate) error {

	if err := r.db.Model(&Models.UserModel{}).Create(user).Error; err != nil {
		return err
	}

	return nil

}
