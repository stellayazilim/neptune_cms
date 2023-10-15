package token_entity

type Token struct {
	ID          ID          `db:"id" sql:"id"`
	Value       Value       `db:"value" sql:"value"`
	TokenType   TokenType   `db:"type" sql:"type"`
	TokenStatus TokenStatus `db:"status" sql:"status"`
	AccountID   AccountID   `db:"account_id" sql:"account_id"`
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
type ID uint64
type Value string
type Tokens []*Token
type AccountID string

type TokenType string
type TokenStatus string

// #endregion

// #region token status enum
func NewDefaultTokenStatus() TokenStatus {
	return new(TokenStatus).VALID()
}

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
