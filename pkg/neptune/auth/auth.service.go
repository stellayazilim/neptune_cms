package auth

import (
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/account"
)

type IAuthService interface {
	Signup(dto *SignupDto) error
	Signin(dto *SigninDto) (string, error)
}

type authServiceRepositories struct {
	account account.IAccountRepository
}
type authService struct {
	repositories authServiceRepositories
	helpers      IAuthHelper
}

func AuthService(accountRepository account.IAccountRepository, helpers IAuthHelper) IAuthService {

	repositories := authServiceRepositories{
		account: accountRepository,
	}

	return &authService{
		repositories: repositories,
		helpers:      helpers,
	}
}

func (s *authService) Signup(dto *SignupDto) error {
	return s.repositories.account.CreateAccount(&models.Account{
		Email:    dto.Email,
		Password: s.helpers.HashPassword(dto),
	})
}

func (s *authService) Signin(dto *SigninDto) (string, error) {
	acc := &models.Account{
		Email: dto.Email,
	}
	err := s.repositories.account.GetAccountByEmail(acc)

	if s.helpers.ComparePassword(acc, dto) {
		return "", err
	}

	return "", PasswordDoesNotMatchError
}
