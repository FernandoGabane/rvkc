package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func RoleNotFoundError(ctx *gin.Context) {
	erro 			 := "role_not_found"
	description 	 := "Regra não encontrada."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusBadRequest)
}