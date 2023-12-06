package Models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"primarykey"`
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	CreatedAt time.Time `gorm:"index autoCreateTime:unix"`
	UpdatedAt time.Time `gorm:"autoCreateTime:unix"`
}

func (*UserModel) TableName() string {
	return "user"
}
