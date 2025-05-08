package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func ClubStartAtError(ctx *gin.Context) {
	erro 	     := "club_start_at"
	description  := "A data de inicio do club não pode ser menor que o horário atual."
	defaultError := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultError) }

	NewDefaultError(ctx, errorResponse, http.StatusConflict)
}