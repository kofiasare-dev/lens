package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() fiber.StructValidator {
	return &Validator{validate: validator.New()}
}

func (v *Validator) Validate(out any) error {
	return v.validate.Struct(out)
}

func FormatValidationError(err error) (errs []string) {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			errs = append(errs, fmt.Sprintf(
				"Field '%s' failed validation on the '%s' rule",
				fieldErr.Field(),
				fieldErr.Tag(),
			))
		}
	}

	return errs
}
