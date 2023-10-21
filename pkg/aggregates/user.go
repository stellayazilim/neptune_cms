package aggregates

import (
	"github.com/stellayazilim/neptune_cms/pkg/entities"
)

type User struct {
	account *entities.Account
	profile *entities.Profile
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

func (a *User) GetProfile() *entities.Profile {
	return a.profile
}

func (a *User) SetProfile(profile entities.Profile) {
	a.profile = &profile
}
