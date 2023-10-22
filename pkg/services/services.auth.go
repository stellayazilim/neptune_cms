package services

import (
	"os"
	"time"

	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	domain_auth "github.com/stellayazilim/neptune_cms/pkg/domain/domain.auth"
	domain_user "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	domain_user_mem "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user/memory"
	"github.com/stellayazilim/neptune_cms/pkg/entities"
	"github.com/stellayazilim/neptune_cms/pkg/libs/bcrypt"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

type IAuthService interface {
	Login(dto domain_auth.LoginRequest) (domain_auth.LoginResponse, error)
	Register(dto domain_auth.RegisterRequest) error
	Refresh(dto domain_auth.LoginRequest) (domain_auth.LoginResponse, error)
}

type AuthService struct {
	Repositories struct {
		User domain_user.IUserRepository
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

func (s *AuthService) Login(request domain_auth.LoginRequest) (domain_auth.LoginResponse, error) {

	user, err := s.Repositories.User.GetByEmail(value_objects.Email(request.Body.Email))
	response := *new(domain_auth.LoginResponse)
	if err != nil {
		return response, err
	}
	// validate password
	if !bcrypt.ComparePassword(user.GetAccount().Password, []byte(request.Body.Password)) {
		return response, domain_auth.ErrPasswordInvalid
	}
	// Create payload
	payload := domain_auth.CreatePasetoPayload(request.Body.Email, time.Minute*20)
	// generate access token

	accessToken, err := payload.CreatePasetoTokenByPayload([]byte(os.Getenv("PASETO_ACCESS_SYMMETRIC_KEY")))

	if err != nil {
		return response, err
	}

	response.AccessToken = accessToken
	// generate refresh token
	refreshToken, err := payload.CreatePasetoTokenByPayload([]byte(os.Getenv("PASETO_REFRESH_SYMMETRIC_KEY")))

	if err != nil {
		return response, err
	}

	response.RefreshToken = refreshToken
	return response, nil
}
func (s *AuthService) Register(dto domain_auth.RegisterRequest) error {

	acc := entities.NewAccount()

	hash, err := bcrypt.GenHash(value_objects.Password(dto.Body.Password))

	if err != nil {
		return err
	}

	acc.Password = hash
	acc.Email = value_objects.Email(dto.Body.Email)

	user := aggregates.NewUser()
	user.SetAccount(*acc)

	if err := s.Repositories.User.Create(user); err != nil {

		return err
	}

	return nil
}

func (s *AuthService) Refresh(request domain_auth.LoginRequest) (domain_auth.LoginResponse, error) {

	response := *new(domain_auth.LoginResponse)

	// Create payload
	payload := domain_auth.CreatePasetoPayload(request.Body.Email, time.Minute*20)
	// generate access token

	accessToken, err := payload.CreatePasetoTokenByPayload([]byte(os.Getenv("PASETO_ACCESS_SYMMETRIC_KEY")))

	if err != nil {
		return response, err
	}

	response.AccessToken = accessToken
	// generate refresh token
	refreshToken, err := payload.CreatePasetoTokenByPayload([]byte(os.Getenv("PASETO_REFRESH_SYMMETRIC_KEY")))

	if err != nil {
		return response, err
	}

	response.RefreshToken = refreshToken
	return response, nil

}

func AuthServiceWithMemUserRepository() ServiceConfig[AuthService] {
	return func(us *AuthService) error {
		userRepository := domain_user_mem.New()
		us.Repositories.User = userRepository
		return nil
	}
}
