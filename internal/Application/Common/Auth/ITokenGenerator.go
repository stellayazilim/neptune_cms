package Auth

type ITokenGenerator interface {
	Generate(TokenPayload) string
}
