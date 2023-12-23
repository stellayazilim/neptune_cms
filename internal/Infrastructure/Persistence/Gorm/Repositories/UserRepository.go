package Repositories

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stellayazilim/neptune.application/Repositories"
	DomainModels "github.com/stellayazilim/neptune.domain/Common/Models"
	"github.com/stellayazilim/neptune.domain/User"
	. "github.com/stellayazilim/neptune.domain/User"
	UserEntities "github.com/stellayazilim/neptune.domain/User/Entities"
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

	roles := make([]*Models.RoleModel, 0)

	for _, ro := range user.Roles {
		roles = append(roles, &Models.RoleModel{
			ID:    ro.ID.GetValue(),
			Name:  ro.Name,
			Perms: ro.Perms.GetValue(),
		})
	}

	if err := r.db.Model(&Models.UserModel{}).Create(&Models.UserModel{
		ID:        user.Root.ID.GetValue(),
		FirstName: user.Root.FirstName,
		LastName:  user.Root.LastName,
		Password:  user.Root.Password,
		Email:     user.Root.Email,
		Roles:     roles,
	}).Error; err != nil {
		switch err.(*pgconn.PgError).Code {
		case "23505":
			return User.ErrUserAlreadyExistException
		}

		return err
	}
	return nil
}

func (r *userRepository) FindByEmail(email string) (User.UserAggregate, error) {
	user := &Models.UserModel{}

	if err := r.db.Model(user).Preload("Roles").Find(user, "email=?", email).Error; err != nil {
		return User.EmptyUserAggregate(), nil
	}

	roles := make([]*UserEntities.RoleEntity, 0)

	for _, ro := range user.Roles {
		roles = append(roles, &UserEntities.RoleEntity{
			Name:  ro.Name,
			Perms: ValueObjects.NewPerms(ro.Perms),
		})
	}
	return UserAggregate{
		AggregateRoot: DomainModels.AggregateRoot[UserEntities.UserEntity]{
			Root: UserEntities.UserEntity{
				Entity: DomainModels.Entity[ValueObjects.UserID]{
					ID: ValueObjects.ParseUserID(user.ID),
				},
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Password:  user.Password,
			},
		},

		Roles: roles,
	}, nil
}

func (r *userRepository) UpdatePasswordByEmail(email string, password []byte) error {

	r.db.Model(&Models.UserModel{}).Where(&Models.UserModel{
		Email: email,
	}).Updates(&Models.UserModel{
		Password: password,
	})
	return nil
}

func (r *userRepository) Update(data User.UserAggregate) error {

	user := &Models.UserModel{
		ID:        data.Root.ID.GetValue(),
		FirstName: data.Root.FirstName,
		LastName:  data.Root.LastName,
		Password:  data.Root.Password,
		Email:     data.Root.Email,
	}

	r.db.Model(user).Save(user)
	return nil
}
