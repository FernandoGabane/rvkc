package dto

import (
	"time"
	"github.com/go-playground/validator/v10"
)

func init() {
	Validate.RegisterValidation("time_format", TimeFormatValidator)
	Validate.RegisterValidation("date_format", DateFormatValidator)
	Validate.RegisterStructValidation(TimeConflictValidator, ClubRequest{})
}


type ClubRequest struct {
	Name       *string `json:"name"       validate:"required,min=4,max=50"`
	Date	   *string `json:"date"       validate:"required,date_format"`
	StartAt    *string `json:"start_at"   validate:"required,time_format"`
	EndAt      *string `json:"end_at"     validate:"required,time_format"`
	Slots	   *uint   `json:"slots"      validate:"required"`
}


func TimeFormatValidator(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	return len(timeStr) == 5 && timeStr[2] == ':' // Exemplo básico: "HH:MM"
}


func DateFormatValidator(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse("2006-01-02", dateStr) // formato yyyy-MM-dd
	return err == nil
}

func TimeConflictValidator(sl validator.StructLevel) {
	req := sl.Current().Interface().(ClubRequest)

	if req.StartAt == nil || req.EndAt == nil {
		return
	}

	layout := "15:04"

	start, err1 := time.Parse(layout, *req.StartAt)
	end, err2   := time.Parse(layout, *req.EndAt)

	if err1 == nil && err2 == nil {
		if !start.Before(end) {
			sl.ReportError(req.EndAt, "end_at", "EndAt", "conflict_date", "")
		}
	}
}
