package aggregate_test

import (
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	"github.com/stellayazilim/neptune_cms/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserAggregate(t *testing.T) {

	user := aggregates.NewUser()

	a := assert.New(t)

	a.Equal(user.GetAccount(), entities.NewAccount())
}
