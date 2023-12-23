package Models

import (
	"time"

	"github.com/google/uuid"
)

type RoleModel struct {
	ID        uuid.UUID `gorm:"primarykey"`
	Name      string    `gorm:"unique"`
	Perms     []byte
	CreatedAt time.Time `gorm:"index autoCreateTime:unix"`
	UpdatedAt time.Time `gorm:"autoCreateTime:unix"`
}

func (*RoleModel) TableName() string {
	return "role"
}
