package dto

import (
    "regexp"

    "github.com/go-playground/validator/v10"
)

var Validate = validator.New()

var telefoneRegex = regexp.MustCompile(`^[0-9]{8,15}$`)

func validarTelefone(fl validator.FieldLevel) bool {
    return telefoneRegex.MatchString(fl.Field().String())
}

func init() {
    Validate.RegisterValidation("telefone_numeric", validarTelefone)
}

type PilotRequest struct {
    Document *string `json:"document" validate:"required,len=11,numeric"`  
    Name     *string `json:"name" validate:"required_without=Update,omitempty,min=3,max=100"`      
    Phone    *string `json:"phone" validate:"required_without=Update,omitempty,telefone_numeric"` 
    Email    *string `json:"email" validate:"required_without=Update,omitempty,email"`            
}