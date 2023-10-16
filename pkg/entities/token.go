package entities

import "github.com/google/uuid"

type Token struct {
	BaseEntity,
	ID uuid.NullUUID `db:"id"`
	Value       string      `db:"value"`
	TokenType   TokenType   `db:"type"`
	TokenStatus TokenStatus `db:"status"`
	AccountID   string      `db:"account_id"`
}

// #region token type enum
const (
	TokenType_AUTHENTICATION TokenType = "AUTHENTICATION"
	TokenType_VALIDATION     TokenType = "VALIDATION"
	TokenType_OTP            TokenType = "OTP"
	TokenType_SECRET         TokenType = "SECRET"
	TokenType_API_KEY        TokenType = "API_KEY"
	TokenType_REFRESH        TokenType = "REFRESH"
)

// #endregion

// #region token value objects
type Tokens []*Token

type TokenType string
type TokenStatus string

// #endregion

// #region token status enum

func NewEmptyTokenStatus() TokenStatus {
	return *new(TokenStatus)
}

func (t *TokenStatus) VALID() TokenStatus {
	return TokenStatus("VALID")
}

func (t *TokenStatus) INVALID() TokenStatus {
	return TokenStatus("INVALID")
}

// #endregion
