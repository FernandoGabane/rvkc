package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func AccountNotFoundError(ctx *gin.Context) {
	erro 			 := "account_not_found"
	description 	 := "Conta não encontrada."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusBadRequest)
}