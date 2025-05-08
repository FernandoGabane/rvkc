package context_error

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


func ConfirmationStatusAlreadyRegisteredError(ctx *gin.Context, status string) {
	descriptionError := fmt.Sprintf("Usuário já está %v no club.", translateStatus(status))
	
	erro 	     := "confirmation_status_already_registered"
	description  := descriptionError
	defaultError := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultError) }

	NewDefaultError(ctx, errorResponse, http.StatusConflict)
}


func translateStatus(status string) string {
	switch status {
		case "CONFIRMED":
			return "confirmado"
		case "UNCONFIRMED":
			return "desconfirmado"
		default:
			return "pendente"	
	}
}