package auth_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
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

func (t *TSAuthHelpers) TestCreateToken() {

	h := auth.AuthHelper()
	a := &models.Account{
		Email: "jhon@doe.com",
	}
	k := utils.GenNCharString(32)
	d := time.Minute * 10
	token := h.CreateToken([]byte(k), a, d)
	t.Run("It should create paseto payload", func() {

		t.NotEmpty(token)
		t.True(strings.HasPrefix(token, "v2.local."))
	})

	t.Run("It should decrypt to payload", func() {
		n := time.Now()
		p := &auth.PasetoPayload{}
		h.Paseto().Decrypt(token, []byte(k), p, nil)

		t.Equal(p.Audience, "Account")
		t.Equal(p.Issuer, "api.neptunecms.com")
		t.Equal(p.Subject, "jhon@doe.com")
		t.LessOrEqual(p.IssuedAt, n)
		t.LessOrEqual(p.NotBefore, n)
		t.LessOrEqual(p.Expiration, n.Add(d))
	})

}
