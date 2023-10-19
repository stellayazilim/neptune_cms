package value_objects_test

import (
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {

	type testCase struct {
		test        string
		expectedErr error
		password    string
	}

	tCases := []testCase{
		{
			test:     "should convert password into []byte",
			password: "1234",
		},
		{
			test:     "should create empty []byte",
			password: "",
		},
	}

	for _, tCase := range tCases {

		t.Run(tCase.test, func(t *testing.T) {

			p := value_objects.NewPassword(tCase.password)

			assert.Equal(t, tCase.password, p.ToString())

		})
	}

}
