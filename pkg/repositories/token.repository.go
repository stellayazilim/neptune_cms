package repositories

import (
	"context"

	token_entity "github.com/stellayazilim/neptune_cms/pkg/entities/token.entity"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
)

type ITokenRepository interface {
	Create(context.Context, *token_entity.Token) error
	Find(context.Context) (*token_entity.Tokens, error)
	FindByTokenId(context.Context, *token_entity.ID) (*token_entity.Token, error)
	FindByAccountId(context.Context, *token_entity.AccountID) (*token_entity.Token, error)
	Update(context.Context, *token_entity.Token) error
	Delete(context.Context, *token_entity.ID) error
}
type tokenRepository struct {
	postgres *postgres.Postgres
}

func TokenRepository(p *postgres.Postgres) ITokenRepository {
	return &tokenRepository{
		postgres: p,
	}
}

// #region Create token
func (r *tokenRepository) Create(ctx context.Context, token *token_entity.Token) error {

	tx, err := r.postgres.DB.BeginTxx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	result, err := tx.ExecContext(ctx,
		/* sql */ `
			INSERT INTO Tokens (value, type, status, account_id) VALUES ($1, $2, $3, $4)
		`, token.Value, token.TokenType, token.TokenStatus, token.AccountID)

	if err != nil {

		tx.Rollback()
		return err
	}

	if _, err := result.RowsAffected(); err != nil {

		return err
	}
	tx.Commit()
	return nil
}

// #endregion

// #region find all tokens
func (r *tokenRepository) Find(ctx context.Context) (*token_entity.Tokens, error) {
	t := new(token_entity.Tokens)

	err := r.postgres.DB.SelectContext(ctx, t, "SELECT * FROM Tokens")

	if err != nil {
		return t, err
	}
	return t, nil
}

// #endregion

// #region Find token by tokenId
func (r *tokenRepository) FindByTokenId(ctx context.Context, id *token_entity.ID) (*token_entity.Token, error) {
	t := &token_entity.Token{}

	if err := r.postgres.DB.GetContext(ctx, t,
		/* sql */ `
			SELECT * FROM Tokens WHERE id=$1`, id,
	); err != nil {

		return t, err
	}
	return t, nil
}

// #endregion

// #region find token by account id
func (r *tokenRepository) FindByAccountId(ctx context.Context, id *token_entity.AccountID) (*token_entity.Token, error) {

	t := &token_entity.Token{}
	if err := r.postgres.DB.GetContext(ctx, t,
		/* sql */ `
			SELECT * FROM Tokens WHERE id=$1`, id,
	); err != nil {

		return t, err
	}

	return t, nil
}

func (r *tokenRepository) Update(ctx context.Context, t *token_entity.Token) error {

	tx, err := r.postgres.DB.BeginTxx(ctx, nil)

	_, err = tx.ExecContext(ctx,
		/* sql */ `
	UPDATE Tokens SET status=$2 WHERE id=$1`, t.ID, t.TokenStatus)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *tokenRepository) Delete(ctx context.Context, id *token_entity.ID) error {
	return nil
}
