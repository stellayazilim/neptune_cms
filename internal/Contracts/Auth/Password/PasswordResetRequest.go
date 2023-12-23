package PasswordContract

type PasswordResetRequestBody struct {
	NewPassword string `json:"newPassword" validate:"required"`
}
