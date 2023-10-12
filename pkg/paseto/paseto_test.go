package paseto

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
	"github.com/stretchr/testify/suite"
)

type TSPaseto struct {
	suite.Suite
}

func TestTSPaseto(t *testing.T) {
	suite.Run(t, new(TSPaseto))
}

func (s *TSPaseto) TestCreatePasetoPayload() {

	acc := &models.Account{
		Email:    "jhon@doe.com",
		Password: []byte("1234"),
	}
	duration := time.Second * 20
	payload := CreatePasetoPayload(acc, duration)

	s.Run("payload must not nil", func() {

		s.Assertions.NotNil(payload)
	})
	s.Run("payload must have Audience as Account", func() {

		s.Assertions.Equal(payload.Audience, "Account")
	})
	s.Run("payload must have Issuer as Stella", func() {

		s.Assertions.Equal(payload.Issuer, "Stella")
		s.Assertions.LessOrEqual(payload.IssuedAt, time.Now())
		s.Assertions.LessOrEqual(payload.NotBefore, time.Now())
	})

	s.Run("payload must have IssuedAt and NotBefore fields and exact same values", func() {

		s.Assertions.LessOrEqual(payload.IssuedAt, time.Now())
		s.Assertions.LessOrEqual(payload.NotBefore, time.Now())
		s.Assertions.Equal(payload.IssuedAt, payload.NotBefore)
	})

	s.Run("Payload subject must match account email", func() {

		s.Assertions.Equal(payload.Subject, acc.Email)
	})

	s.Run("Payload must contain Jti as type of uuid", func() {

		s.Assertions.IsType(uuid.UUID{}, payload.Jti)
	})

}

func (s *TSPaseto) TestCreatePasetoTokenFromPayload() {
	acc := &models.Account{
		Email:    "jhon@doe.com",
		Password: []byte("1234"),
	}
	duration := time.Second * 20
	payload := CreatePasetoPayload(acc, duration)
	key := utils.GenNCharString(32)

	s.Run("Should create token", func() {
		token, _ := payload.CreatePasetoTokenByPayload([]byte(key))
		s.Assertions.NotNil(token)
	})

	s.Run("Token must be prefixed `v2.local.`", func() {
		token, _ := payload.CreatePasetoTokenByPayload([]byte(key))
		s.Assertions.True(strings.HasPrefix(token, "v2.local."))
	})

	s.Run("Token must be decrypted to payload", func() {
		token, _ := payload.CreatePasetoTokenByPayload([]byte(key))
		decrypted := new(PasetoPayload)
		err := paseto.NewV2().Decrypt(token, []byte(key), decrypted, nil)

		s.Assertions.Nil(err)
		s.Assertions.Equal(decrypted.Issuer, payload.Issuer)
	})

	s.Run("Must not throw error", func() {

		_, err := payload.CreatePasetoTokenByPayload([]byte(key))

		s.Assertions.Nil(err)
	})

	s.Run("Must return error on key less then 32 char", func() {

		skey := utils.GenNCharString(12)
		_, err := payload.CreatePasetoTokenByPayload([]byte(skey))

		s.Assertions.NotNil(err)
	})

}
