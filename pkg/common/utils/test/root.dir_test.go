package utils_test

import (
	"strings"
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetRootDir(t *testing.T) {

	t.Run("Should retrive root dir", func(t *testing.T) {

		root := utils.GetRootDir()
		assert.True(t, strings.HasSuffix(root, "/neptune_cms"))
	})
}
