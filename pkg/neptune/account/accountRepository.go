package account

import (
	"fmt"

	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
)

type IAccountRepository interface {
	GetAccountByEmail(a *models.Account) error
	CreateAccount(a *models.Account) error
}
type accountRepository struct {
	postgres *postgres.Postgres
}

func AccountRepository(p *postgres.Postgres) IAccountRepository {
	return &accountRepository{
		postgres: p,
	}
}

func (r *accountRepository) GetAccountByEmail(a *models.Account) error {
	fmt.Println(a.Email)
	var err error
	tx := r.postgres.DB.MustBegin()

	if err != nil {
		fmt.Println(err)
	}

	err = tx.Get(a,

		/* sql */
		`SELECT id, email, password FROM Accounts WHERE email=$1`, &a.Email)
	return err
}

func (r *accountRepository) CreateAccount(a *models.Account) error {
	return nil
}
