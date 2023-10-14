package tokentype

type TokenType string

const (
	AUTHENTICATION           = "AUTHENTICATION"
	VALIDATION               = "VALIDATION"
	OTP                      = "OTP"
	SECRET                   = "SECRET"
	API_KEY                  = "API_KEY"
	REFRESH        TokenType = "REFRESH"
)
