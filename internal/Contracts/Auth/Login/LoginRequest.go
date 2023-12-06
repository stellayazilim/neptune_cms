package AuthContract

type LoginRequestHeader struct {
	XTokenLookup [2]string `json:"x-token-lookup"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
