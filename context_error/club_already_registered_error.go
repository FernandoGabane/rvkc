package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func ClubAlreadyRegisteredError(ctx *gin.Context) {
	erro 	     := "club_already_registered"
	description  := "Já existe um club cadastrado neste mesmo horário."
	defaultError := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultError) }

	NewDefaultError(ctx, errorResponse, http.StatusConflict)
}