package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/entities"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
)

type ITokenRepository interface {
	Create(context.Context, *entities.Token) error
	Find(context.Context) (*entities.Tokens, error)
	FindByTokenId(context.Context, uint64) (*entities.Token, error)
	FindByAccountId(context.Context, uuid.UUID) (*entities.Token, error)
	Update(context.Context, *entities.Token) error
	Delete(context.Context, uint64) error
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
func (r *tokenRepository) Create(ctx context.Context, token *entities.Token) error {

	tx, err := r.postgres.DB.BeginTxx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("acc_id :", token.AccountID)
	result, err := tx.ExecContext(ctx,
		/* sql */ `
			INSERT INTO Tokens (value, type, status, account_id) VALUES ($1, $2, $3, $4)
		`, token.Value, token.TokenType, token.TokenStatus, token.AccountID)

	if err != nil {
		fmt.Println("error happ:", err.Error())
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
func (r *tokenRepository) Find(ctx context.Context) (*entities.Tokens, error) {
	t := new(entities.Tokens)

	err := r.postgres.DB.SelectContext(ctx, t, "SELECT * FROM Tokens")

	if err != nil {
		return t, err
	}
	return t, nil
}

// #endregion

// #region Find token by tokenId
func (r *tokenRepository) FindByTokenId(ctx context.Context, id uint64) (*entities.Token, error) {
	t := &entities.Token{}

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
func (r *tokenRepository) FindByAccountId(ctx context.Context, id uuid.UUID) (*entities.Token, error) {

	t := &entities.Token{}
	if err := r.postgres.DB.GetContext(ctx, t,
		/* sql */ `
			SELECT * FROM Tokens WHERE id=$1`, id,
	); err != nil {

		return t, err
	}

	return t, nil
}

func (r *tokenRepository) Update(ctx context.Context, t *entities.Token) error {

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

func (r *tokenRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}
