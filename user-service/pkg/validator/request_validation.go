package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateRequest(request interface{}) error {
	var errorMessages []string
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", err.Field()))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is email", err.Field()))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters log", err.Field(), err.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is max", err.Field()))
			}
		}
		return errors.New("Validasi gagal: " + joinMessage(errorMessages))
	}
	return nil
}

func joinMessage(errorMessages []string) string {
	result := ""
	for i, message := range errorMessages {
		if i > 0 {
			result += " , " + message
		}
		result += "  " + message
	}
	return result
}
