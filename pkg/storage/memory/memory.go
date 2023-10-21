package memory

import (
	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
)

var Users *map[uuid.UUID]*aggregates.User

func InitMemoryUser() {
	u := make(map[uuid.UUID]*aggregates.User, 0)
	Users = &u
}
