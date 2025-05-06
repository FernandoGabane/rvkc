package middleware

import (
	"regexp"
	"strconv"
	"github.com/go-playground/validator/v10"
)

var cpfOnlyDigits = regexp.MustCompile(`[^0-9]`)


func DocumentValidator(fl validator.FieldLevel) bool {
	cpf := cpfOnlyDigits.ReplaceAllString(fl.Field().String(), "")

	if len(cpf) != 11 {
		return false
	}

	allSame := true
	for i := 1; i < 11 && allSame; i++ {
		if cpf[i] != cpf[0] {
			allSame = false
		}
	}
	if allSame {
		return false
	}

	for t := 9; t < 11; t++ {
		sum := 0
		for i := 0; i < t; i++ {
			num, _ := strconv.Atoi(string(cpf[i]))
			sum += num * (t + 1 - i)
		}
		d := (sum * 10) % 11
		if d == 10 {
			d = 0
		}
		expected, _ := strconv.Atoi(string(cpf[t]))
		if d != expected {
			return false
		}
	}

	return true
}
