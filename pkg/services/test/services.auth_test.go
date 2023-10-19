package services_test

import (
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/services"
	"github.com/stretchr/testify/suite"
)

type TestCase struct {
	test           string
	input          any
	expectedResult any
	expectedError  error
}
type TSAuthService struct {
	suite.Suite
}

func (t *TSAuthService) TestAuthServiceFactory() {

	testCases := []TestCase{
		{
			test:          "It should construct Auth service",
			input:         nil,
			expectedError: nil,
		},
		{
			test:          "It should construct Auth service with User memory repository",
			input:         services.AuthServiceWithMemUserRepository,
			expectedError: nil,
		},
	}

	for _, tCase := range testCases {

		t.Run(tCase.test, func() {
			var authService services.IAuthService
			var err error
			if tCase.input == nil {
				authService, err = services.AuthServiceFactory()

			} else {
				authService, err = services.AuthServiceFactory(tCase.input.(func(*services.AuthService) error))
			}

			t.Implements((*services.IAuthService)(nil), authService)
			t.Nil(err)

		})
	}

}

func TestTSAuthService(t *testing.T) {
	suite.Run(t, new(TSAuthService))
}
