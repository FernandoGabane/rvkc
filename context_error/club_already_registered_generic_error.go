package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func ClubAlreadyRegisteredGenericError(ctx *gin.Context) {
	erro 	     := "club_already_registered_generic"
	description  := "Erro ao verificar se já existe um club cadastrado neste mesmo horário."
	defaultError := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultError) }

	NewDefaultError(ctx, errorResponse, http.StatusInternalServerError)
}