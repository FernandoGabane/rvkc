package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func RolePersistError(ctx *gin.Context) {
	erro 			 := "internal_server_error"
	description 	 := "Erro ao salvar role."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusInternalServerError)
}