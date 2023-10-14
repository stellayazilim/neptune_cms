package models

import "github.com/google/uuid"

type Token struct {
	ID        uint64        `db:"id" sql:"id"`
	Token     string        `db:"token" sql:"token"`
	TokenType string        `db:"token" sql:"token"`
	Account   *Account      `db:"account" sql:"account"`
	AccountID uuid.NullUUID `db:"account_id" sql:"account_id"`
}
