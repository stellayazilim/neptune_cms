package Repositories

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune.application/Repositories"
	UserEntities "github.com/stellayazilim/neptune.domain/User/Entities"
	"github.com/stellayazilim/neptune.infrastructure/Persistence/Gorm/Models"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

// Create implements Repositories.IRoleRepository.
func (r *roleRepository) Create(role UserEntities.RoleEntity) error {

	rl := Models.RoleModel{
		ID:    role.ID.GetValue(),
		Name:  role.Name,
		Perms: role.Perms.GetValue(),
	}

	if err := r.db.Save(&rl).Error; err != nil {
		return err
	}

	return nil
}

// GetAll implements Repositories.IRoleRepository.
func (*roleRepository) GetAll() ([]UserEntities.RoleEntity, error) {
	panic("unimplemented")
}

// GetById implements Repositories.IRoleRepository.
func (*roleRepository) GetById(uuid.UUID) (UserEntities.RoleEntity, error) {
	panic("unimplemented")
}

// GetByUserId implements Repositories.IRoleRepository.
func (*roleRepository) GetByUserId(uuid.UUID) ([]UserEntities.RoleEntity, error) {
	panic("unimplemented")
}

// Update implements Repositories.IRoleRepository.
func (*roleRepository) Update(UserEntities.RoleEntity) error {
	panic("unimplemented")
}

func GormRoleRepository(db *gorm.DB) Repositories.IRoleRepository {
	return &roleRepository{
		db: db,
	}
}
