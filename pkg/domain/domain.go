package domain

type RepositoryConfig[c any] func(repository *c) error

type RequestGetUserById struct {
	Params struct {
		ID string `param:"id"`
	}
}

type UserResponse struct {
	Body struct {
		Email string `json:"email"`
	}
}
