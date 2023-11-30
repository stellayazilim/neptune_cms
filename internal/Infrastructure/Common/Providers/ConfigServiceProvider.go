package Providers

import "time"

type configServiceProvider struct {
	PasetoAccessSymmetricKey  [32]rune
	PasetoRefreshSymmetricKey [32]rune
	PasetoIssuer              string
	PasetoAudience            string
	PasetoAccessExpiration    time.Duration
	PasetoRefreshExpiration   time.Duration
}

func ConfigServiceProvider() *configServiceProvider {
	return &configServiceProvider{}
}
