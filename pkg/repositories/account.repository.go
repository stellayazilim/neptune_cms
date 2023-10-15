package repositories

import (
	"context"
	"fmt"

	account_entity "github.com/stellayazilim/neptune_cms/pkg/entities/account.entity"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
)

type IAccountRepository interface {
	Find(context.Context) (*account_entity.Accounts, error)
	FindById(context.Context, *account_entity.ID) (*account_entity.Account, error)
	FindByEmail(context.Context, *account_entity.Email) (*account_entity.Account, error)
	Create(context.Context, *account_entity.Account) error
	Update(context.Context, *account_entity.Account) error
	Delete(context.Context, *account_entity.Account) error
}
type accountRepository struct {
	postgres *postgres.Postgres
}

func AccountRepository(p *postgres.Postgres) IAccountRepository {
	return &accountRepository{
		postgres: p,
	}
}

func (r *accountRepository) Find(ctx context.Context) (*account_entity.Accounts, error) {

	a := new(account_entity.Accounts)

	return a, nil
}

func (r *accountRepository) FindById(ctx context.Context, id *account_entity.ID) (*account_entity.Account, error) {
	a := new(account_entity.Account)

	return a, nil
}

func (r *accountRepository) FindByEmail(ctx context.Context, email *account_entity.Email) (*account_entity.Account, error) {
	a := &account_entity.Account{}

	tx := r.postgres.DB.MustBeginTx(ctx, nil)

	err := tx.GetContext(ctx, a,
		/* SQL */ `
		SELECT id, email, password FROM Accounts WHERE email=$1
	`, email)

	if err != nil {

		fmt.Println(err)
		tx.Rollback()
		return a, err
	}

	return a, nil
}

func (r *accountRepository) Create(ctx context.Context, account *account_entity.Account) error {
	return nil
}

func (r *accountRepository) Update(ctx context.Context, account *account_entity.Account) error {
	return nil
}

func (r *accountRepository) Delete(ctx context.Context, account *account_entity.Account) error {
	return nil
}
