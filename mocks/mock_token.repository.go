package mocks

import (
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockTokenRepository struct {
	mock.Mock
}

func (m *MockTokenRepository) CreateRefreshToken(account *models.Account) (*models.Token, error) {
	args := m.Called(account)

	return args.Get(0).(*models.Token), args.Error(1)
}
