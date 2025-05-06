package dto

import (
	"rvkc/middleware"
	"rvkc/models"
)


const (
    ADMIN   = "ADMIN"
    COACH   = "COACH"
    DEFAULT = "DEFAULT"
)


func init() {
    middleware.Validate.RegisterValidation("phone_numeric_format", middleware.PhoneValidator)
    middleware.Validate.RegisterValidation("document_validator", middleware.DocumentValidator)
}


type AccountRequest struct {
    Document       *string               `json:"document" validate:"required,len=11,numeric,document_validator"`  
    Name           *string               `json:"name"     validate:"required_without=Update,omitempty,min=3,max=100"`      
    Phone          *string               `json:"phone"    validate:"required_without=Update,omitempty,phone_numeric_format"` 
    Email          *string               `json:"email"    validate:"required_without=Update,omitempty,email,max=100"`
    Roles          []*models.Role        `json:"roles"`            
}