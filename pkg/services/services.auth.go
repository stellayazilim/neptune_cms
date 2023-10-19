package services

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	domain_user_mem "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
	"github.com/stellayazilim/neptune_cms/pkg/entities"
	"github.com/stellayazilim/neptune_cms/pkg/libs/bcrypt"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

type IAuthService interface {
	Login(dto domain_auth.LoginDto) ([]string, error)
	Register(dto domain_auth.RegisterDto) error
}

type AuthService struct {
	Repositories struct {
		User domain_user.IAccountRepository
	}
}

func AuthServiceFactory(cfgs ...ServiceConfig[AuthService]) (IAuthService, error) {

	as := new(AuthService)

	for _, cfg := range cfgs {
		if err := cfg(as); err != nil {
			return as, err
		}
	}
	return as, nil
}

func (s *AuthService) Login(dto domain_auth.LoginDto) ([]string, error) {

	user, err := s.Repositories.User.GetByEmail(value_objects.Email(dto.Email))

	if err != nil {
		return []string{}, err
	}
	// validate password
	if !bcrypt.ComparePassword(user.GetAccount().Password, dto.Password) {
		return []string{}, domain_auth.PasswordInvalidErr
	}
	// Create payload
	payload := domain_auth.CreatePasetoPayload(dto.Email, time.Minute*20)
	// generate access token
	accessToken, err := payload.CreatePasetoTokenByPayload([]byte(os.Getenv("PASETO_ACCESS_SYMMETRIC_KEY")))

	if err != nil {
		return []string{}, err
	}
	// generate refresh token
	refreshToken, err := payload.CreatePasetoTokenByPayload([]byte(os.Getenv("PASETO_REFRESH_SYMMETRIC_KEY")))

	if err != nil {
		return []string{}, err
	}

	return []string{
		accessToken,
		refreshToken,
	}, nil
}
func (s *AuthService) Register(dto domain_auth.RegisterDto) error {

	acc := entities.NewAccount()

	hash, err := bcrypt.GenHash(value_objects.Password(dto.Password))

	if err != nil {
		return err
	}

	acc.Password = hash
	acc.Email = value_objects.Email(hash)

	user := aggregates.NewUser()
	user.SetAccount(*acc)

	return s.Repositories.User.Create(user)
}

func AuthServiceWithMemUserRepository(s *AuthService) error {
	users := make(map[uuid.UUID]aggregates.User)

	userRepository := domain_user_mem.New(users)
	s.Repositories.User = userRepository
	return nil
}
