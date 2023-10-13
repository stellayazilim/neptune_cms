package auth_test

import (
	"testing"

	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stretchr/testify/suite"
)

type TSAuthHelpers struct {
	suite.Suite
}

func TestAuthHelpers(t *testing.T) {
	suite.Run(t, new(TSAuthHelpers))
}

func (t *TSAuthHelpers) TestHashPassword() {

	dto := &auth.SignupDto{
		Password: "1234",
	}

	helper := auth.AuthHelper()

	hashedPw := helper.HashPassword(dto)

	t.Run("hashed password must be not null", func() {
		t.NotEmpty(hashedPw)
		t.NotNil(hashedPw)
	})

	t.Run("hashed password must be comparable non hashed one", func() {

		t.True(helper.ComparePassword(&models.Account{
			Password: hashedPw,
		}, &auth.SigninDto{
			Password: "1234",
		}))
	})
}
