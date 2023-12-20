package PasswordContract

type PasswordResetRequestBody struct {
	OldPassword []byte    `json:"oldPassword" validate:"string"`
	NewPassword [2][]byte `json:"newPassword" validate:"required,dive,required"`
}
