package tokentype

type TokenType string

const (
	VALIDATION     TokenType = "VALIDATION"
	AUTHENTICATION TokenType = "AUTHENTICATION"
	OTP            TokenType = "OTP"
	SECRET         TokenType = "SECRET"
	API_KEY        TokenType = "API_KEY"
	REFRESH        TokenType = "REFRESH"
)
