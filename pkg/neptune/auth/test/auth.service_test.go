package auth_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stellayazilim/neptune_cms/mocks"
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stretchr/testify/suite"
)

type TSAuthService struct {
	suite.Suite
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(TSAuthService))
}

func (t *TSAuthService) TestSignup() {
	mockAccountRepo := mocks.MockAccountRepository{}
	authService := auth.AuthService(&mockAccountRepo)

	t.Run("should create account", func() {

		adto := &auth.SignupDto{
			Email:    "jhon@doe.com",
			Password: "1234",
		}

		am := &models.Account{
			Email:    adto.Email,
			Password: adto.HashPassword(),
		}
		mockAccountRepo.On("CreateAccount", am).Return(nil)
		err := authService.Signup(adto)

		t.Nil(err)
	})

	t.Run("should get account by email", func() {
		var err error
		var token string
		adto := &auth.SigninDto{
			Email:    "jhon@doe.com",
			Password: "1234",
		}

		mockAccountRepo.On("GetAccountByEmail", &models.Account{
			Email: adto.Email,
		}).Return(nil)

		t.Nil(err)
		token, err = authService.Signin(adto)

		t.Nil(err)
		t.Equal(token, "")

	})
}
