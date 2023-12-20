package RegisterContract

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stellayazilim/neptune.contracts/Validator"
)

type RegisterRequestBody struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  []byte `json:"password" validate:"required"`
}

func (r *RegisterRequestBody) Validate(ctx *fiber.Ctx) *Validator.IValidationErrorResponse {
	v := validator.New()
	ve := make([]Validator.IValidationError, 0)
	errs := v.Struct(r)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		ve = append(ve,
			Validator.IValidationError{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
				Error:       true,
			},
		)
	}

	return &Validator.IValidationErrorResponse{
		Type:     "Validation error",
		Instance: string(ctx.Request().URI().Path()),
		Detail:   &ve,
		Status:   422,
	}

}
