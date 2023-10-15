package token_test

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/token"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type TSTokenRepository struct {
	suite.Suite
	token token.ITokenRepository
	mock  sqlmock.Sqlmock
}

func TestTokenRepository(t *testing.T) {

	suite.Run(t, new(TSTokenRepository))

}

func (t *TSTokenRepository) SetupTestSuite() {
	db, mock, err := sqlmock.New()
	t.Nil(err)
	dbx := sqlx.NewDb(db, "postgres")

	pg := &postgres.Postgres{
		DB: dbx,
	}
	tokenRepo := token.TokenRepository(pg)
	t.token = tokenRepo
	t.mock = mock

}

func (t *TSTokenRepository) TestCreateRefreshToken() {

	t.Run("It should create refresh token", func() {
		fmt.Println(t.token)
	})

}
