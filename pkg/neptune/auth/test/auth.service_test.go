package auth_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stellayazilim/neptune_cms/mocks"
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TSAuthService struct {
	suite.Suite
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(TSAuthService))
}

func (t *TSAuthService) TestSignup() {
	mockAccountRepo := new(mocks.MockAccountRepository)
	mockAuthHelper := new(mocks.MockAuthHelper)
	authService := auth.AuthService(mockAccountRepo, mockAuthHelper)

	t.Run("should signup", func() {

		adto := &auth.SignupDto{
			Email:    "jhon@doe.com",
			Password: "1234",
		}

		ac := &models.Account{
			Email:    adto.Email,
			Password: []byte("$2a$10$NDCBnYIfoCPk/n6HJKJLFexxQPdOS528F62iwznU2nkFiiPS3siBq"),
		}

		mockAuthHelper.On("HashPassword", adto).Return([]byte("$2a$10$NDCBnYIfoCPk/n6HJKJLFexxQPdOS528F62iwznU2nkFiiPS3siBq"))

		mockAccountRepo.On("CreateAccount", ac).Return(nil)
		err := authService.Signup(adto)
		t.Nil(err)
	})

	t.Run("should signin by email", func() {
		// var err error
		// var token string

		adto := &auth.SigninDto{
			Email:    "jhon@doe.com",
			Password: "1234",
		}

		var ac *models.Account = &models.Account{
			Email: adto.Email,
		}
		mockAccountRepo.On("GetAccountByEmail", ac).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(0).(*models.Account)
				arg.Password = []byte("$2a$10$NDCBnYIfoCPk/n6HJKJLFexxQPdOS528F62iwznU2nkFiiPS3siBq")

			})

		mockAuthHelper.On("ComparePassword", ac, adto).Return(true)
		// t.Nil(err)
		token, err := authService.Signin(adto)

		t.Nil(err)
		t.Equal(token, "")

	})
}
