package main

import (
	"os"
	"testing"

	"github.com/stellayazilim/neptune_cms/internal/rest"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	t.Run("Should inject .env files", func(t *testing.T) {
		// load env
		utils.InjectEnv()

		assert.Equal(t, os.Getenv("POSTGRES_DB"), "neptune")
	})

	t.Run("Should create new Rest app", func(t *testing.T) {
		r := rest.Rest()
		assert.NotNil(t, r)
		assert.Implements(t, ((*rest.IRest)(nil)), r.(rest.IRest))

	})

}
