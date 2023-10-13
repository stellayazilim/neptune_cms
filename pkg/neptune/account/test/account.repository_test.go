package account_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/account"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type TSAccountRepository struct {
	suite.Suite
}

func TestAccount(t *testing.T) {
	suite.Run(t, new(TSAccountRepository))
}

func (t *TSAccountRepository) TestAccountRepository() {

	mockDB, mock, _ := sqlmock.New()

	defer mockDB.Close()
	ps := &postgres.Postgres{
		DB: sqlx.NewDb(mockDB, "postgres"),
	}

	t.Run("Should get account by email", func() {

		ar := account.AccountRepository(ps)
		dto := &auth.SigninDto{
			Email:    "jhon@doe.com",
			Password: "1234",
		}

		acc := &models.Account{
			Email: dto.Email,
			ID:    uuid.NewString(),
		}

		accountMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"}).
			AddRow(acc.ID, acc.Email, acc.Password)
		mock.ExpectBegin()
		mock.ExpectQuery(
			"SELECT id, email, password FROM Accounts WHERE email=?",
		).WithArgs(dto.Email).WillReturnRows(accountMockRow)

		err := ar.GetAccountByEmail(acc)

		t.Nil(err)
	})
}
