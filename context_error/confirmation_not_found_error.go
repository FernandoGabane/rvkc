package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func ConfirmationNotFoundError(ctx *gin.Context) {
	erro 			 := "confirmation_not_found"
	description 	 := "Cconfirmação não encontrada."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusBadRequest)
}