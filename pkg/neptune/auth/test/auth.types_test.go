package auth_test

import (
	"encoding/json"
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stretchr/testify/assert"
)

func TestPasetoPayload(t *testing.T) {
	p := &auth.PasetoPayload{
		Audience: "Test",
		Issuer:   "neptune_test",
		Subject:  "jhon@doe.com",
	}
	t.Run("Should marshall into []byte", func(t *testing.T) {

		decoded, err := json.Marshal(p)

		assert.Nil(t, err)
		assert.NotEmpty(t, decoded)

	})

	t.Run("Should unmarsall into PasetoPayload", func(t *testing.T) {
		decoded, err := json.Marshal(p)

		assert.Nil(t, err)
		exp := new(auth.PasetoPayload)

		err = json.Unmarshal(decoded, exp)

		assert.Nil(t, err)

		assert.Equal(t, exp, p)
	})
}
