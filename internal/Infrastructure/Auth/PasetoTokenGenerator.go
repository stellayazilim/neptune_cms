package Auth

import (
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Auth"
)

type pasetoTokenProvider struct {
	configService string
}

func (p *pasetoTokenProvider) Generate(paylod TokenPayload) string {

	return ""
}

func PasetoTokenProvider() ITokenGenerator {
	return &pasetoTokenProvider{}
}
