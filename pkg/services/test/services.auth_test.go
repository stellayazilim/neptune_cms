package services_test

import (
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/common/utils"
	domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	domain_user_mem "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
	"github.com/stellayazilim/neptune_cms/pkg/services"
	"github.com/stellayazilim/neptune_cms/pkg/storage/memory"
	"github.com/stretchr/testify/suite"
)

type TSAuthService struct {
	suite.Suite
}

func (t *TSAuthService) TestAuthServiceFactory() {

	t.Run("Should create service with mem user repository", func() {
		as, err := services.AuthServiceFactory(services.AuthServiceWithMemUserRepository())
		t.Implements((*services.IAuthService)(nil), as)
		t.Nil(err)
	})

	t.Run("Should throw nil pointer exception if not a repository provided", func() {
		as, _ := services.AuthServiceFactory()
		s := as.(*services.AuthService)
		t.Equal(s.Repositories.User, nil)

	})
}

func (t *TSAuthService) TestRegister() {
	tcases := []struct {
		test      string
		input     domain_auth.RegisterRequest
		expectErr error
	}{
		{
			test: "Should create new user",
			input: func() domain_auth.RegisterRequest {

				d := domain_auth.RegisterRequest{}
				d.Body.Email = "jhon@doe.com"
				d.Body.Password = "1234"
				return d
			}(),
			expectErr: nil,
		},
		{
			test: "Should return user already exist err",
			input: func() domain_auth.RegisterRequest {

				d := domain_auth.RegisterRequest{}
				d.Body.Email = "jhon@doe.com"
				d.Body.Password = "1234"
				return d
			}(),
			expectErr: domain_user.UserAreadyExistsError,
		},
	}

	for _, c := range tcases {
		t.Run(c.test, func() {
			as, _ := services.AuthServiceFactory(services.AuthServiceWithMemUserRepository())

			err := as.Register(c.input)

			t.Equal(c.expectErr, err)

		})
	}
}

func (t *TSAuthService) TestLogin() {

	tcases := []struct {
		test        string
		input       domain_auth.LoginRequest
		expectErr   error
		expectedRes domain_auth.LoginResponse
	}{
		{
			test: "should get tokens",
			input: func() domain_auth.LoginRequest {
				d := domain_auth.LoginRequest{}
				d.Body.Email = "jhon@doe.com"
				d.Body.Password = "1234"
				return d
			}(),
			expectErr:   nil,
			expectedRes: domain_auth.LoginResponse{},
		},
		{
			test: "should get user not found error",
			input: func() domain_auth.LoginRequest {
				d := domain_auth.LoginRequest{}
				d.Body.Email = "jhon1@doe.com"
				d.Body.Password = "1234"
				return d
			}(),
			expectErr: domain_user.UserNotFoundError,
		},
	}

	as, _ := services.AuthServiceFactory(services.AuthServiceWithMemUserRepository())

	d := domain_auth.RegisterRequest{}
	d.Body.Email = "jhon@doe.com"
	d.Body.Password = "1234"

	as.Register(d)
	for _, c := range tcases {

		t.Run(c.test, func() {

			tokens, err := as.Login(c.input)
			t.Equal(c.expectErr, err)

			if err == nil {
				t.NotEmpty(tokens)
			} else {
				t.Empty(tokens)
			}

		})
	}

}

func (t *TSAuthService) TestAuthSericeConfigs() {
	tcases := []struct {
		test      string
		input     services.ServiceConfig[services.AuthService]
		expectErr error
		expect    domain_user.IUserRepository
	}{
		{
			test:      "should create auth service with memory repository",
			input:     services.AuthServiceWithMemUserRepository(),
			expectErr: nil,
			expect:    domain_user_mem.New(),
		},
	}

	for _, c := range tcases {

		t.Run(c.test, func() {
			s, err := services.AuthServiceFactory(c.input)
			t.Nil(err)
			t.Equal(c.expect, s.(*services.AuthService).Repositories.User)
		})
	}
}

func (t *TSAuthService) SetupSuite() {
	utils.InjectEnv(utils.GetRootDir() + "/env/.env.test")

}

func (t *TSAuthService) SetupTest() {
	memory.Users = nil
	memory.InitMemoryUser()
}
func TestTSAuthService(t *testing.T) {
	suite.Run(t, new(TSAuthService))
}
