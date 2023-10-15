package mocks

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
	DB *sqlx.DB
}
