package repositories

import (
	"context"

	entitiy_errors "github.com/stellayazilim/neptune_cms/pkg/common/errors"
	"github.com/stellayazilim/neptune_cms/pkg/entities"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
)

type IAccountRepository interface {
	Find(context.Context) (*entities.Accounts, error)
	FindById(context.Context, string) (*entities.Account, error)
	FindByEmail(context.Context, string) (*entities.Account, error)
	Create(context.Context, *entities.Account) error
	Update(context.Context, *entities.Account) error
	Delete(context.Context, *entities.Account) error
}
type accountRepository struct {
	postgres *postgres.Postgres
}

func AccountRepository(p *postgres.Postgres) IAccountRepository {
	return &accountRepository{
		postgres: p,
	}
}

func (r *accountRepository) Find(ctx context.Context) (*entities.Accounts, error) {

	a := new(entities.Accounts)

	if err := r.postgres.DB.SelectContext(ctx, a,
		/* sql */ `
		SELECT id, email, password FROM Accounts
	`,
	); err != nil {
		return a, err
	}
	return a, nil
}

func (r *accountRepository) FindById(ctx context.Context, id string) (*entities.Account, error) {
	a := &entities.Account{}

	if err := r.postgres.DB.GetContext(ctx, a,
		/* sql */ `
		SELECT id, email, password FROM Accounts WHERE id=$1
	`, id); err != nil {
		return a, entitiy_errors.RECORD_NOT_FOUND_ERROR
	}
	return a, nil
}

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (*entities.Account, error) {
	a := &entities.Account{}
	if err := r.postgres.DB.GetContext(ctx, a,
		/* SQL */ `
		SELECT id, email, password FROM Accounts WHERE email=$1
	`, email); err != nil {

		return a, entitiy_errors.RECORD_NOT_FOUND_ERROR
	}

	return a, nil
}

func (r *accountRepository) Create(ctx context.Context, account *entities.Account) error {

	tx, _ := r.postgres.DB.BeginTxx(ctx, nil)
	defer tx.Commit()

	row := tx.QueryRowxContext(ctx,
		/* sql */ `INSERT INTO Accounts (email, password) VALUES ( $1, $2 ) RETURNING id`,
		account.Email, account.Password)

	if err := row.Scan(&account.ID); err != nil {

		tx.Rollback()
		return err
	}

	return nil
}

func (r *accountRepository) Update(ctx context.Context, account *entities.Account) error {
	return nil
}

func (r *accountRepository) Delete(ctx context.Context, account *entities.Account) error {
	return nil
}
