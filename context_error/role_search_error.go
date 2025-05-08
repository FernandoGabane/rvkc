package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func RoleSearchError(ctx *gin.Context) {
	erro 			 := "internal_server_error"
	description 	 := "Erro ao buscar role."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusInternalServerError)
}