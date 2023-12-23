package LoginContract

type LoginRequestHeader struct {
	XTokenLookup [2]string `json:"x-token-lookup"`
}

type LoginRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
