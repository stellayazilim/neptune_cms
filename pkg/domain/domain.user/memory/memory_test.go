package memory_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	domain_account "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	"github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
	"github.com/stretchr/testify/suite"
)

type TSMDomainMemoryRepository struct {
	suite.Suite
}

type testCase struct {
	test           string
	input          any
	expectedError  error
	expectedResult any
}

func TestTSDomainMemoryRepository(t *testing.T) {
	suite.Run(t, new(TSMDomainMemoryRepository))
}

func (t *TSMDomainMemoryRepository) TestCreate() {
	accountMemRepository := memory.New(nil)
	happy := aggregates.NewUser()
	happy.GetAccount().ID.UUID = uuid.New()
	happy.GetAccount().ID.Valid = true

	sad := aggregates.NewUser()

	testCases := []testCase{
		{
			test:           "it should create account on valid input",
			input:          happy,
			expectedError:  nil,
			expectedResult: nil,
		},
		{
			test:           "it should throw error on invalid id",
			input:          sad,
			expectedError:  domain_account.UserInvalidIDError,
			expectedResult: nil,
		},
		{
			test:           "it should throw error if account already exists with given id",
			input:          happy,
			expectedError:  domain_account.UserAreadyExistsError,
			expectedResult: nil,
		},
	}
	for _, tCase := range testCases {

		t.Run(tCase.test, func() {
			err := accountMemRepository.Create(tCase.input.(aggregates.User))

			t.Equal(err, tCase.expectedError)
		})
	}
}

func (t *TSMDomainMemoryRepository) TestGetAll() {

	happy := aggregates.NewUser()
	happy.GetAccount().ID.UUID = uuid.New()
	happy.GetAccount().ID.Valid = true

	testCases := []testCase{
		{
			test:           "should return empty slice if account does not exist",
			expectedError:  nil,
			expectedResult: make([]aggregates.User, 0),
			input:          memory.New(make(map[uuid.UUID]aggregates.User)),
		},
		{
			test:          "should return non empty slice if accounts exists",
			expectedError: nil,
			expectedResult: []aggregates.User{
				happy,
			},
			input: memory.New(map[uuid.UUID]aggregates.User{
				happy.GetAccount().ID.UUID: happy,
			}),
		},
	}

	for _, tCase := range testCases {

		t.Run(tCase.test, func() {

			in, _ := tCase.input.(domain_account.IUserRepository)

			res, err := in.GetAll()
			t.Equal(res, tCase.expectedResult.([]aggregates.User))
			t.Equal(err, tCase.expectedError)
		})
	}
}

func (t *TSMDomainMemoryRepository) TestGetById() {

	userId := uuid.New()
	user := aggregates.NewUser()
	user.GetAccount().ID.UUID = userId
	users := make(map[uuid.UUID]aggregates.User, 0)

	users[userId] = user
	mrepo := memory.New(
		users,
	)

	testCases := []testCase{
		{
			test:           "it should user",
			input:          userId,
			expectedError:  nil,
			expectedResult: user,
		},
		{
			test:           "it should return user not found error",
			input:          uuid.New(),
			expectedError:  domain_account.UserNotFoundError,
			expectedResult: aggregates.NewUser(),
		},
	}

	for _, tcase := range testCases {

		t.Run(tcase.test, func() {

			user, err := mrepo.GetById(tcase.input.(uuid.UUID))

			t.Equal(tcase.expectedError, err)
			t.Equal(tcase.expectedResult, user)

		})
	}
}
