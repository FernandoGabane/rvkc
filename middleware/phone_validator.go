package middleware

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)


var telefoneRegex = regexp.MustCompile(`^[0-9]{10,11}$`)


func PhoneValidator(fl validator.FieldLevel) bool {
    return telefoneRegex.MatchString(fl.Field().String())
}