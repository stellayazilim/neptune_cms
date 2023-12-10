package Repositories

import (
	"github.com/stellayazilim/neptune.application/Repositories"
	DomainModels "github.com/stellayazilim/neptune.domain/Common/Models"
	"github.com/stellayazilim/neptune.domain/User"
	. "github.com/stellayazilim/neptune.domain/User"
	"github.com/stellayazilim/neptune.domain/User/Entities"
	"github.com/stellayazilim/neptune.domain/User/ValueObjects"
	"github.com/stellayazilim/neptune.infrastructure/Persistence/Gorm/Models"
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

	if err := r.db.Model(&Models.UserModel{}).Create(&Models.UserModel{
		ID:        user.Root.ID.GetID(),
		FirstName: user.Root.FirstName,
		LastName:  user.Root.LastName,
		Password:  user.Root.Password,
		Email:     user.Root.Email,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (User.UserAggregate, error) {
	user := &Models.UserModel{}

	if err := r.db.Model(user).Find(user, "email=?", email).Error; err != nil {
		return User.EmptyUserAggregate(), nil
	}

	return UserAggregate{
		AggregateRoot: DomainModels.AggregateRoot[Entities.UserEntity]{
			Root: Entities.UserEntity{
				Entity: DomainModels.Entity[ValueObjects.UserID]{
					ID: ValueObjects.ParseUserID(user.ID),
				},
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Password:  user.Password,
			},
		},
	}, nil

}
