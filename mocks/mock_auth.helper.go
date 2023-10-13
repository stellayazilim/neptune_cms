package mocks

import (
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stretchr/testify/mock"
)

type MockAuthHelper struct {
	mock.Mock
}

func (m *MockAuthHelper) ComparePassword(acc *models.Account, dto *auth.SigninDto) bool {
	args := m.Called(acc, dto)
	return args.Bool(0)
}

func (m *MockAuthHelper) HashPassword(dto *auth.SignupDto) []byte {

	args := m.Called(dto)
	return args.Get(0).([]byte)
}
