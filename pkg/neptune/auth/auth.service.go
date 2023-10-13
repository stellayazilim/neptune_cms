package auth

import (
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/account"
)

type IAuthService interface {
	Signup(*SignupDto) error
	Signin(dto *SigninDto) (string, error)
}

type authServiceRepositories struct {
	account account.IAccountRepository
}
type authService struct {
	repositories authServiceRepositories
}

func AuthService(account account.IAccountRepository) IAuthService {

	repositories := authServiceRepositories{
		account: account,
	}

	return &authService{
		repositories: repositories,
	}
}

func (s *authService) Signup(dto *SignupDto) error {
	return s.repositories.account.CreateAccount(&models.Account{
		Email:    dto.Email,
		Password: dto.HashPassword(),
	})
}

func (s *authService) Signin(dto *SigninDto) (string, error) {
	acc := &models.Account{
		Email: dto.Email,
	}
	err := s.repositories.account.GetAccountByEmail(acc)

	return "", err
}
