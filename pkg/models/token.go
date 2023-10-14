package models

import (
	TokenType "github.com/stellayazilim/neptune_cms/pkg/enums/token.type"
)

type Token struct {
	ID        uint64              `db:"id" sql:"id"`
	Token     string              `db:"token" sql:"token"`
	TokenType TokenType.TokenType `db:"token" sql:"token"`
	Account   *Account            `db:"account" sql:"account"`
	AccountID string              `db:"account_id" sql:"account_id"`
}
