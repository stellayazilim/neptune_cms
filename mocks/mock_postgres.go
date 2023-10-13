package mocks

import "github.com/stretchr/testify/mock"

type MockPostgres struct {
	mock.Mock
}

func (m *MockPostgres) Get(dest interface{}, query string, args ...interface{}) error {
	cargs := m.Called(dest, query, args)

	return cargs.Error(0)
}
