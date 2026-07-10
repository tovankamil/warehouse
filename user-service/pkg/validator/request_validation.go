package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validate.value

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
				errorMessages = append(errorMessages, fmt.SprintF("%s is required", err.Field()))
			case "email":
				errorMessages = append(errorMessages, fmt.SprintF("%s is email", err.Field()))
			case "min":
				errorMessages = append(errorMessages, fmt.SprintF("%s is min", err.Field()))
			case "max":
				errorMessages = append(errorMessages, fmt.SprintF("%s is max", err.Field()))
			}
			return errors.New("Validasi gagal " + joinMessage)
		}
		return errorMessages
	}
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
