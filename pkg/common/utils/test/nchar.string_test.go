package utils_test

import (
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestNcharString(t *testing.T) {

	asserts := assert.New(t)
	t.Run("should generate n char string", func(t *testing.T) {
		str := utils.GenNCharString(32)
		asserts.Equal(len(str), 32)
	})
}
