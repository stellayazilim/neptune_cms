package Providers

import (
	"os"
	"strconv"
	"time"
)

type ConfigService struct {
	AccessTokenSymmetricKey  string
	RefreshTokenSymmetricKey string
	TokenPrefix              string
	TokenIssuer              string
	TokenAudience            string
	TokenAccessExpiration    time.Duration
	TokenRefreshExpiration   time.Duration
	PostgresDatabase         string
	PostgresHost             string
	PostgresPort             string
	PostgresUser             string
	PostgresPassword         string
	IsDevolopment            bool
}

func ConfigServiceProvider() *ConfigService {

	accessExp, _ := strconv.ParseInt(os.Getenv("ACCESS_EXPIRATION"), 10, 32)
	refreshExp, _ := strconv.ParseInt(os.Getenv("REFRESH_EXPIRATION"), 10, 32)
	return &ConfigService{
		AccessTokenSymmetricKey:  os.Getenv("ACCESS_TOKEN_SYMMETRIC_KEY"),
		RefreshTokenSymmetricKey: os.Getenv("REFRESH_TOKEN_SYMMETRIC_KEY"),
		TokenAccessExpiration:    time.Duration(time.Minute * time.Duration(accessExp)),
		TokenRefreshExpiration:   time.Duration(time.Minute * time.Duration(refreshExp)),
		TokenIssuer:              os.Getenv("TOKEN_ISSUER"),
		TokenAudience:            os.Getenv("TOKEN_AUDIENCE"),
		TokenPrefix:              os.Getenv("TOKEN_PREFIX"),
		PostgresDatabase:         os.Getenv("POSTGRES_DB"),
		PostgresHost:             os.Getenv("POSTGRES_HOST"),
		PostgresPort:             os.Getenv("POSTGRES_PORT"),
		PostgresUser:             os.Getenv("POSTGRES_USER"),
		PostgresPassword:         os.Getenv("POSTGRES_PASSWORD"),
		IsDevolopment:            os.Getenv("GO_ENV") == "development",
	}
}
