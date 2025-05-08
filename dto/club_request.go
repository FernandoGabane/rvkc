package dto

import (
	"rvkc/middleware"

	"github.com/go-playground/validator/v10"
)


func init() {
	middleware.Validate.RegisterValidation("custom_time_format", middleware.CustomTimeFormatValidator)
	middleware.Validate.RegisterStructValidation(TimeConflictValidator, ClubRequest{})
}


type ClubRequest struct {
	Name      *string     		   	 `json:"name"       validate:"required,min=4,max=50"`
	StartAt   *middleware.CustomTime `json:"start_at"   validate:"required,custom_time_format"`
	EndAt  	  *middleware.CustomTime `json:"end_at"     validate:"required,custom_time_format"`
	AccountId *string     		   	 `json:"account_id" validate:"required"`
	Slots     *uint       		   	 `json:"slots"      validate:"required"`
}


func TimeConflictValidator(sl validator.StructLevel) {
	req := sl.Current().Interface().(ClubRequest)

	if req.StartAt == nil || req.EndAt == nil {
		return
	}

	if !req.StartAt.Time.Before(req.EndAt.Time) {
		sl.ReportError(req.EndAt, "end_at", "EndAt", "conflict_date", "")
	}
}
