package main_test

import (
	"os"
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/utils/config"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	t.Run("it should inject env variables without error", func(t *testing.T) {

		err := config.InjectEnv()

		assert.Nil(t, err)

		assert.Equal(t, "neptune", os.Getenv("POSTGRES_DB"))

	})
}
