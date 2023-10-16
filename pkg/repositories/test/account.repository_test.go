package repositories_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	entitiy_errors "github.com/stellayazilim/neptune_cms/pkg/common/errors"
	"github.com/stellayazilim/neptune_cms/pkg/entities"
	"github.com/stellayazilim/neptune_cms/pkg/repositories"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type TSAccountRepository struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	mockDB *sql.DB
	r      repositories.IAccountRepository
}

func (s *TSAccountRepository) SetupSubTest() {

	mockDB, mock, _ := sqlmock.New()

	s.mockDB = mockDB
	ps := postgres.Postgres{
		DB: sqlx.NewDb(mockDB, "sqlmock"),
	}
	s.mock = mock
	s.r = repositories.AccountRepository(&ps)
}
func TestAccountRepository(t *testing.T) {
	utils.InjectEnv(utils.GetRootDir() + "/env/.env.test")
	suite.Run(t, new(TSAccountRepository))
}

func (t *TSAccountRepository) TestFindByEmail() {
	defer t.mockDB.Close()
	expected := &entities.Account{
		ID:       uuid.New().String(),
		Email:    "jhon@doe.com",
		Password: []byte("$2a$10$xQkLtRJEYODBGkUFYeiFPOeBCKM0nkVeQ/GwDRkMeOMaFZagVdbcy"),
	}

	t.Run("It should query [FROM Accounts WHERE email=<email>]", func() {
		acMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"}).
			AddRow(expected.ID, expected.Email, expected.Password)

		t.mock.
			ExpectQuery("SELECT id, email, password FROM Accounts WHERE email=?").
			WithArgs(&expected.Email).
			WillReturnRows(acMockRow)

		_, err := t.r.FindByEmail(context.TODO(), expected.Email)
		t.Nil(err)

	})

	t.Run("It should get account", func() {
		acMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"}).
			AddRow(expected.ID, expected.Email, expected.Password)

		t.mock.
			ExpectQuery("SELECT id, email, password FROM Accounts WHERE email=?").
			WithArgs(expected.Email).
			WillReturnRows(acMockRow)

		acc, err := t.r.FindByEmail(context.TODO(), expected.Email)

		t.Equal(expected, acc)
		t.Nil(err)
	})

	t.Run("It should return error if account not found", func() {
		acMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"})
		t.mock.
			ExpectQuery("SELECT id, email, password FROM Accounts WHERE email=?").
			WithArgs(expected.Email).
			WillReturnRows(acMockRow)

		acc, err := t.r.FindByEmail(context.TODO(), expected.Email)

		t.Equal(&entities.Account{}, acc)
		t.NotNil(err)

		t.ErrorIs(err, entitiy_errors.RECORD_NOT_FOUND_ERROR)

	})
}

func (t *TSAccountRepository) TestFind() {
	defer t.mockDB.Close()
	expected := &entities.Accounts{
		{
			ID:       uuid.New().String(),
			Email:    "jhon@doe.com",
			Password: []byte("$2a$10$xQkLtRJEYODBGkUFYeiFPOeBCKM0nkVeQ/GwDRkMeOMaFZagVdbcy1"),
		},
		{
			ID:       uuid.New().String(),
			Email:    "jhon1@doe.com",
			Password: []byte("$2a$10$xQkLtRJEYODBGkUFYeiFPOeBCKM0nkVeQ/GwDRkMeOMaFZagVdbcy1"),
		},
		{
			ID:       uuid.New().String(),
			Email:    "jhon2@doe.com",
			Password: []byte("$2a$10$xQkLtRJEYODBGkUFYeiFPOeBCKM0nkVeQ/GwDRkMeOMaFZagVdbcy1"),
		},
	}

	t.Run("It should run query [FROM Accounts]", func() {
		acMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"})

		for _, row := range *expected {
			acMockRow.AddRow(row.ID, row.Email, row.Password)
		}

		t.mock.
			ExpectQuery("SELECT id, email, password FROM Accounts").
			WillReturnRows(acMockRow)

		_, err := t.r.Find(context.TODO())

		t.Nil(err)

	})

	t.Run("It should return empty slice if not record found", func() {
		acMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"})

		t.mock.
			ExpectQuery("SELECT id, email, password FROM Accounts").
			WillReturnRows(acMockRow)

		accs, err := t.r.Find(context.TODO())

		t.Equal(new(entities.Accounts), accs)
		t.Nil(err)
	})
}

func (t *TSAccountRepository) TestFindById() {
	defer t.mockDB.Close()
	expected := &entities.Account{
		ID:       uuid.New().String(),
		Email:    "jhon@doe.com",
		Password: []byte("$2a$10$xQkLtRJEYODBGkUFYeiFPOeBCKM0nkVeQ/GwDRkMeOMaFZagVdbcy"),
	}
	t.Run("It should run query [FROM Accounts WHERE id = ?]", func() {

		acMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"}).
			AddRow(expected.ID, expected.Email, expected.Password)

		t.mock.ExpectQuery("SELECT id, email, password FROM Accounts WHERE id=?").
			WithArgs(expected.ID).
			WillReturnRows(acMockRow)

		_, err := t.r.FindById(context.TODO(), expected.ID)

		t.Nil(err)
	})

	t.Run("It should get account", func() {

		acMockRow := sqlmock.
			NewRows([]string{"id", "email", "password"}).
			AddRow(expected.ID, expected.Email, expected.Password)

		t.mock.ExpectQuery("SELECT id, email, password FROM Accounts WHERE id=?").
			WithArgs(expected.ID).
			WillReturnRows(acMockRow)

		_, err := t.r.FindById(context.TODO(), expected.ID)

		t.Nil(err)
	})
}

func (t *TSAccountRepository) TestCreate() {

	t.Run("It should run query [INSERT INTO Accounts]", func() {
		defer t.mockDB.Close()
		expected := &entities.Account{
			ID:       uuid.New().String(),
			Email:    "jhon@doe.com",
			Password: []byte("$2a$10$xQkLtRJEYODBGkUFYeiFPOeBCKM0nkVeQ/GwDRkMeOMaFZagVdbcy"),
		}
		acMockRow := sqlmock.
			NewRows([]string{"id"}).
			AddRow(expected.ID)
		t.mock.ExpectBegin()

		t.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO Accounts (email, password) VALUES ( $1, $2 ) RETURNING id`)).WithArgs(
			expected.Email, expected.Password).WillReturnRows(acMockRow)

		err := t.r.Create(context.TODO(), &entities.Account{
			Email:    expected.Email,
			Password: expected.Password,
		})
		t.Nil(err)
		t.mock.ExpectCommit()

	})

	t.Run("it should populte account with inserted id", func() {

		expected := &entities.Account{
			ID:       uuid.New().String(),
			Email:    "jhon@doe.com",
			Password: []byte("$2a$10$xQkLtRJEYODBGkUFYeiFPOeBCKM0nkVeQ/GwDRkMeOMaFZagVdbcy"),
		}
		acMockRow := sqlmock.
			NewRows([]string{"id"}).
			AddRow(expected.ID)
		t.mock.ExpectBegin()

		t.mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO Accounts (email, password) 
				VALUES ( $1, $2 ) RETURNING id`)).
			WithArgs(
				expected.Email, expected.Password).WillReturnRows(acMockRow)

		t.r.Create(context.TODO(), &entities.Account{
			Email:    expected.Email,
			Password: expected.Password,
		})

		t.mock.ExpectCommit()
	})
}
