package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func JsonReadError(ctx *gin.Context) {
	erro 			 := "json_read"
	description 	 := "Erro ao ler o JSON."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusInternalServerError)
}