package entities

type Account struct {
	BaseEntity
	ID       string `db:"id" sql:"id"`
	Email    string `db:"email" sql:"email"`
	Password []byte `db:"password" sql:"password"`
}
type Accounts []*Account

/* value objects */
