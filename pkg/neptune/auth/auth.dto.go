package auth

type SignupDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *SignupDto) HashPassword() []byte {
	// hash password here
	return []byte("")
}

type SigninDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
