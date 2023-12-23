package Repositories

import (
	"github.com/google/uuid"
	UserEntities "github.com/stellayazilim/neptune.domain/User/Entities"
)

type IRoleRepository interface {
	Create(UserEntities.RoleEntity) error
	GetAll() ([]UserEntities.RoleEntity, error)
	GetByUserId(uuid.UUID) ([]UserEntities.RoleEntity, error)
	GetById(uuid.UUID) (UserEntities.RoleEntity, error)
	Update(UserEntities.RoleEntity) error
}
