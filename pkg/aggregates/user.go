package aggregates

import (
	"github.com/stellayazilim/neptune_cms/pkg/entities"
)

type User struct {
	account   *entities.Account
	profile   *entities.Profile
	addresses *[]*entities.Address
	tokens    *[]*entities.Token
}

func NewUser() User {

	return User{
		account: entities.NewAccount(),
	}
}

func (a *User) GetAccount() *entities.Account {
	return a.account
}

func (a *User) SetAccount(acc entities.Account) {
	a.account = &acc
}

func (a *User) GetAddresses() *[]*entities.Address {
	return a.addresses
}

func (a *User) SetAddresses(addresses []*entities.Address) {
	a.addresses = &addresses
}

func (a *User) GetProfile() *entities.Profile {
	return a.profile
}

func (a *User) SetProfile(profile entities.Profile) {
	a.profile = &profile
}

func (a *User) GetTokens() *[]*entities.Token {
	return a.tokens
}

func (a *User) SetTokens(tokens []*entities.Token) {
	a.tokens = &tokens
}
