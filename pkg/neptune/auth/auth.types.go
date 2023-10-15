package auth

type SignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInTokenSerializer [2]string
