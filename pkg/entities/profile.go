package entities

import "github.com/google/uuid"

type Profile struct {
	Base
	ID uuid.UUID
}
