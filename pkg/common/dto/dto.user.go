package dto

import "github.com/stellayazilim/neptune_cms/pkg/value_objects"

type UserResponseBody struct {
	Email value_objects.Email `json:"email"`
}

type UserResponse struct {
	Body UserResponseBody
}

type UsersResponseBody struct {
	Data         []UserResponseBody `json:"data"`
	Count        uint64             `json:"count"`
	Displaying   uint8              `json:"displaying"`
	CurrentPage  uint8              `json:"currentPage"`
	TotalPage    uint32             `json:"totalPage"`
	NextPage     string             `json:"nextPage"`
	PreviousPage string             `json:"previousPage"`
}
type UsersResponse struct {
	Body UsersResponseBody
}
