package dto

type LoginRequest struct {
	Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterRequest struct {
	Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}
