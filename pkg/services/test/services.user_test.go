package services_test

import (
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/common/utils"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	domain_user_mem "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
	"github.com/stellayazilim/neptune_cms/pkg/services"
	"github.com/stellayazilim/neptune_cms/pkg/storage/memory"
	"github.com/stretchr/testify/suite"
)

type TSUserService struct {
	suite.Suite
}

func (t *TSUserService) SetupSuite() {
	utils.InjectEnv(utils.GetRootDir() + "/env/.env.test")
	memory.InitMemoryUser()
}
func TestUserService(t *testing.T) {

	suite.Run(t, new(TSAuthService))
}

func (t *TSUserService) TestUserServiceFactory() {

	type tCase struct {
		test           string
		expectedErr    error
		input          services.ServiceConfig[services.UserService]
		expectedResult services.IUserService
	}

	tCases := []tCase{
		{
			test:        "It should instantiate userService",
			expectedErr: nil,
			input:       services.UserServiceWithMemUserRepository(),
			expectedResult: &services.UserService{
				Repositories: struct{ User domain_user.IUserRepository }{
					User: domain_user_mem.New(),
				},
			},
		},
	}

	for _, c := range tCases {
		t.Run(c.test, func() {
			s, err := services.UserServiceFactory(services.UserServiceWithMemUserRepository())

			t.Equal(c.expectedErr, err)
			t.Equal(c.expectedResult, s)
		})
	}
}

func (t *TSUserService) TestCreate() {

	type tCase struct {
		test           string
		expextedErr    error
		expectedResult domain_user.UsersResponse
	}

	tCases := []tCase{
		{
			test:           "it should get users",
			expextedErr:    nil,
			expectedResult: domain_user.UsersResponse{},
		},
	}

	for _, c := range tCases {

		t.Run(c.test, func() {

			us, err := services.UserServiceFactory(services.UserServiceWithMemUserRepository())

			t.Nil(err)

			u, err := us.GetAll()
			t.Nil(err)
			t.Equal(c.expectedResult, u)
		})
	}
}

func (t *TSUserService) TestGetById() {

}
