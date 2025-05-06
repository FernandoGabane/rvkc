package context_error

import (
	"github.com/gin-gonic/gin"
)


type DefaultError struct {
	Error 		  *string `json:"error"`
	ParameterName *string `json:"parameter_name"`
	Description   *string `json:"description"`
}


type ErrorResponse struct {
	ErroResponse []DefaultError `json:"error_response"`
}


func NewDefaultError(ctx *gin.Context, errorReponse ErrorResponse, code int) {
	ctx.JSON(code, errorReponse)
}