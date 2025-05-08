package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func JsonInvalidError(ctx *gin.Context) {
	erro 			 := "json_invalid"
	description 	 := "JSON inválido."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusBadRequest)
}