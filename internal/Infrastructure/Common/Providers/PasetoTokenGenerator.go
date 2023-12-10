package Providers

import (
	"log"

	"github.com/o1egl/paseto"
	"github.com/stellayazilim/neptune.application/Common/Interfaces"
	"github.com/stellayazilim/neptune.domain/Auth"
)

type pasetoTokenProvider struct {
	configService *ConfigService
}

func (p *pasetoTokenProvider) GenerateAccessToken(paylod Auth.AccessTokenPayload) string {

	token, err := paseto.NewV2().Encrypt([]byte(p.configService.AccessTokenSymmetricKey), paylod, nil)
	if err != nil {
		// todo handle errors
		log.Fatal(err)
	}
	return token
}

func (p *pasetoTokenProvider) GenerateRefreshToken(paylod Auth.RefreshTokenPayload) string {

	token, err := paseto.NewV2().Encrypt([]byte(p.configService.RefreshTokenSymmetricKey), paylod, nil)
	if err != nil {
		// todo handle errors
		log.Fatal(err)
	}
	return token
}

func PasetoTokenProvider(
	config *ConfigService,
) Interfaces.ITokenProvider {
	return &pasetoTokenProvider{
		configService: config,
	}
}
