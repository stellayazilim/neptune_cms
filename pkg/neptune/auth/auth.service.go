package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/account"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

type IAuthService interface {
	Signup(dto *SignupDto) error
	Signin(dto *SigninDto) ([2]string, error)
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
	fmt.Println(string(utils.GenNCharString(32)))

	return s.repositories.account.CreateAccount(&models.Account{
		Email:    dto.Email,
		Password: s.helpers.HashPassword(dto),
	})
}

func (s *authService) Signin(dto *SigninDto) ([2]string, error) {

	var tokens [2]string
	acc := &models.Account{
		Email: dto.Email,
	}
	if err := s.repositories.account.GetAccountByEmail(acc); err != nil {
		return tokens, err
	}
	if s.helpers.ComparePassword(acc, dto) {
		tokens[0] = s.helpers.CreateToken([]byte(os.Getenv("PASETO_ACCESS_SYMMETRIC_KEY")), acc, time.Minute*20)
		return tokens, nil
	}

	return tokens, PasswordDoesNotMatchError
}
