package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsValidPhone(fl validator.FieldLevel) bool {
	pattern := `^0[0-9]{9}$`
	phoneRegex := regexp.MustCompile(pattern)
	return phoneRegex.MatchString(fl.Field().String())
}
