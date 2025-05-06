package context_error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func ClubInProgressOrFinishedError(ctx *gin.Context) {
	erro 			 := "club_in_progress_or_finished"
	description 	 := "O club está em andamento ou já terminou."
	defaultException := DefaultError{
		Error:          &erro,
		ParameterName:  nil,
		Description:    &description,
	}

	errorResponse := ErrorResponse{ append(make([]DefaultError, 0), defaultException) }
	NewDefaultError(ctx, errorResponse, http.StatusBadRequest)
}