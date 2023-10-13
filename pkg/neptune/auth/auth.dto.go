package auth

type SignupDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
