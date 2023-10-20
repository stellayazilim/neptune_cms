package domain_user

type UserResponse struct {
	Email string `json:"email"`
}

type UsersResponse []UserResponse
