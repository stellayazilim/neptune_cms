package entities

import "time"

type BaseEntity struct {
	Deleted   time.Time `db:"deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
