package helpers

import (
	"github.com/go-playground/validator/v10"
)

// Check checks for validation error
func Check(s interface{}) error {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
