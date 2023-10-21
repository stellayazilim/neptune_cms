package utils_test

import (
	"os"
	"testing"

	utils "github.com/stellayazilim/neptune_cms/pkg/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	asserts := assert.New(t)

	t.Run("should load env variables from .env file", func(t *testing.T) {

		err := utils.InjectEnv(utils.GetRootDir() + "/env/.env.test")
		asserts.NotEmpty(os.Getenv("POSTGRES_DB"))
		asserts.Nil(err)
	})
}
