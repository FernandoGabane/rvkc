package middleware

import (
	"time"
	"github.com/go-playground/validator/v10"
)


func DateFormatValidator(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}


func CustomTimeFormatValidator(fl validator.FieldLevel) bool {
	customTime, ok := fl.Field().Interface().(CustomTime)
	if !ok {
		return true
	}
	return customTime.ParseError == nil
}