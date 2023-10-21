package entities

import "github.com/google/uuid"

type Profile struct {
	Base
	ID uuid.NullUUID
}

func NewProfile() *Profile {

	return &Profile{}
}
