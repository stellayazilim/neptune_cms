package mocks

import (
	"github.com/jmoiron/sqlx"
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
	DB *sqlx.DB
}

func (m *MockAccountRepository) GetAccountByEmail(a *models.Account) error {

	args := m.Called(a)
	return args.Error(0)
}

func (m *MockAccountRepository) CreateAccount(a *models.Account) error {

	args := m.Called(a)

	return args.Error(0)
}
