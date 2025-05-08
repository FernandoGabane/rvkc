package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func ClubNotFoundError(ctx *gin.Context) {
	erro 			 := "club_not_found"
	description 	 := "Club não encontrado."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusBadRequest)
}