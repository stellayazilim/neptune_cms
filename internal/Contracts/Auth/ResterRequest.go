package AuthContract

type RegisterRequestBody struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  []byte `json:"password" validate:"required"`
}
