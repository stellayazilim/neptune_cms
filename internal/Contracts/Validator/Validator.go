package Validator

type (
	IValidationError struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	IValidationErrorResponse struct {
		Type     string              `json:"type"`
		Instance string              `json:"instance"`
		Detail   *[]IValidationError `json:"detail"`
		Status   uint                `json:"status"`
	}
)
