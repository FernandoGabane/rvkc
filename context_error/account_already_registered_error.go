package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func AccountAlreadyRegisteredError(ctx *gin.Context) {
	erro 	     := "account_already_registered"
	description  := "Já existe uma conta cadastrada com este documento."
	defaultError := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultError) }

	NewDefaultError(ctx, errorResponse, http.StatusConflict)
}