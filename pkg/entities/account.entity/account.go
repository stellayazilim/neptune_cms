package account_entity

type Account struct {
	ID       ID       `db:"id"`
	Email    Email    `db:"email"`
	Password Password `db:"password"`
}
type Accounts []Account

/* value objects */
type ID []uint8
type Email string

type Password []byte

func (p *Password) ToString() string {
	return string(*p)
}

func NewAccount() *Account {

	return &Account{
		ID:       ID{},
		Email:    *new(Email),
		Password: Password{},
	}
}
