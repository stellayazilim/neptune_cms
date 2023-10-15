package token

import (
	tokentype "github.com/stellayazilim/neptune_cms/pkg/enums/token.type"
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

type ITokenRepository interface {
	CreateRefreshToken(account *models.Account) (*models.Token, error)
}
type tokenRepository struct {
	postgres *postgres.Postgres
}

func TokenRepository(p *postgres.Postgres) ITokenRepository {
	return &tokenRepository{
		postgres: p,
	}
}

func (r *tokenRepository) GetTokens() *[]*models.Token {

	tokens := make([]*models.Token, 0)

	return &tokens

}

func (r *tokenRepository) CreateRefreshToken(account *models.Account) (*models.Token, error) {

	token := new(models.Token)

	token.AccountID = account.ID
	token.Account = account
	token.Token = utils.Cuid()
	token.TokenType = tokentype.REFRESH

	tx := r.postgres.DB.MustBegin()

	res := tx.QueryRowx(
		/* sql */
		`INSERT INTO Tokens (account_id, token, type) VALUES ($1, $2, $3) RETURNING id`,
		token.AccountID,
		token.Token,
		token.TokenType,
	)

	if err := res.Scan(token.ID); err != nil {
		tx.Rollback()
		return token, err
	}

	tx.Commit()
	return token, nil
}
