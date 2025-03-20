package dto

import (
	"github.com/go-playground/validator/v10"
)
// var V = validator.New()

func init() {
	Validate.RegisterValidation("time_format", TimeFormatValidator)
}


type ClubRequest struct {
	Name       *string `json:"name"       validate:"required,min=4,max=50"`
	Recurrence *string `json:"recurrence" validate:"required,oneof=SEMANAL MENSAL ANUAL"`
	Weekday    *string `json:"weekday"    validate:"required,oneof=SEGUNDA-FEIRA TERÇA-FEIRA QUARTA-FEIRA QUINTA-FEIRA SEXTA-FEIRA SÁBADO DOMINGO"`
	StartAt    *string `json:"start_at"   validate:"required,time_format"`
	EndAt      *string `json:"end_at"     validate:"required,time_format"`
}


func TimeFormatValidator(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	return len(timeStr) == 5 && timeStr[2] == ':' // Exemplo básico: "HH:MM"
}


