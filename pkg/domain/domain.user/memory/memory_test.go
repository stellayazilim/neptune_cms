package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSMDomainUserMemoryRepository struct {
	suite.Suite
}

// type testCase struct {
// 	test           string
// 	input          any
// 	expectedError  error
// 	expectedResult any
// }

func TestTSDomainMemoryRepository(t *testing.T) {
	suite.Run(t, new(TSMDomainUserMemoryRepository))
}

func (t *TSMDomainUserMemoryRepository) TestCreate() {

}

func (t *TSMDomainUserMemoryRepository) TestGetById() {

}
