package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/stellayazilim/neptune_cms/pkg/entities"
	"github.com/stellayazilim/neptune_cms/pkg/libs/paseto"
	"github.com/stellayazilim/neptune_cms/pkg/repositories"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

type IAuthService interface {
	SignIn(context.Context, *SignInDto) (*SignInTokenSerializer, error)
}
type authServiceRepositories struct {
	token   repositories.ITokenRepository
	account repositories.IAccountRepository
}
type authService struct {
	repositories *authServiceRepositories
}

func New(
	accountRepository repositories.IAccountRepository,
	tokenRepository repositories.ITokenRepository,
) IAuthService {
	return &authService{
		repositories: &authServiceRepositories{
			token:   tokenRepository,
			account: accountRepository,
		},
	}
}

func (s *authService) SignIn(ctx context.Context, dto *SignInDto) (*SignInTokenSerializer, error) {

	tokens := new(SignInTokenSerializer)
	account, err := s.repositories.account.FindByEmail(ctx, dto.Email)

	if err != nil {
		return tokens, err
	}

	// create access token
	payload := paseto.CreatePasetoPayload(account.Email, time.Minute*15)
	if accessToken, err := payload.CreatePasetoTokenByPayload([]byte(os.Getenv("PASETO_ACCESS_SYMMETRIC_KEY"))); err != nil {
		fmt.Println("error :", err)
		return tokens, err
	} else {
		tokens[0] = accessToken
		// maybe store token somewhere here
	}

	// create refresh token

	token := &entities.Token{
		Value:       utils.Cuid(),
		TokenType:   entities.TokenType_REFRESH,
		TokenStatus: entities.TokenStatus("VALID"),
		AccountID:   account.ID,
	}

	if err := s.repositories.token.Create(ctx, token); err != nil {
		fmt.Println("error :", err)
		return tokens, err
	}

	tokens[1] = string(token.Value)

	return tokens, nil
}

func (s *authService) SignUp(ctx context.Context, dto *SignUpDto) error {

	// p := account_entity.Password(dto.Password)

	// // hash password
	// if err := bcrypt.GenHash(&p); err != nil {
	// 	return err
	// }

	// // create account entity
	// a := &account_entity.Account{
	// 	Email:    account_entity.Email(dto.Email),
	// 	Password: p,
	// }

	// // persist on database``
	// if err := s.repositories.account.Create(ctx, a); err != nil {
	// 	return err
	// }

	// // if account created successfully create verification token

	// token := &token_entity.Token{
	// 	TokenType: token_entity.TokenType("VERIFICATION"),
	// 	Value:     token_entity.Value(utils.Cuid()),
	// 	AccountID: token_entity.AccountID(a.ID.UUID.String()),
	// }

	// if err := s.repositories.token.Create(token); err != nil {
	// 	return err
	// }

	// send token via email or sms

	return nil
}
