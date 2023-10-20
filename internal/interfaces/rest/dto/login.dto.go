package dto

type LoginDto struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
