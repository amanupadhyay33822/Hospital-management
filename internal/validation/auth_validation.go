package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Auth struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func Validateregister(input Auth) error {
	return validate.Struct(input)
}

func ValidateLogin(input Login) error {
	return validate.Struct(input)
}

func FormatValidationError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				return fmt.Sprintf("%s is required", e.Field())
			case "email":
				return "Invalid email format"
			case "min":
				return fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
			}
		}
	}
	return err.Error()
}
