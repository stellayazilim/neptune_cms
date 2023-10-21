package services_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

//	type TestCase struct {
//		test           string
//		input          any
//		expectedResult any
//		expectedError  error
//	}
type TSAuthService struct {
	suite.Suite
}

func (t *TSAuthService) TestAuthServiceFactory() {

}

func TestTSAuthService(t *testing.T) {
	suite.Run(t, new(TSAuthService))
}
