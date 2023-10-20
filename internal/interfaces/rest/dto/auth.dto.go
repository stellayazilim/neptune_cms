package dto

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
