package models

type Account struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password []byte `db:"password"`
}

type Accounts []Account
