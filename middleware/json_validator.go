package middleware

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"rvkc/context_error"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type CustomValidator struct {
	Validator *validator.Validate
}


func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}


func TranslateValidationError(err error) context_error.ErrorResponse {
	var invalidParameter = context_error.DefaultError{}
	var defaultErrorList []context_error.DefaultError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		for _, e := range validationErrors {
			field := regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(e.Field(), "${1}_${2}")
			field = strings.ToLower(field)
			
			switch e.Tag() {
				case "required":
					invalidParameter = buildInvalidParameterResponse(nil, "O campo " + field + " é obrigatório.", field)
				case "len":
					invalidParameter = buildInvalidParameterResponse(nil, "O campo " + field + " deve ter exatamente " + e.Param() + " caracteres.", field)
				case "min":
					invalidParameter = 	buildInvalidParameterResponse(nil, "O campo " + field + " deve ter no mínimo " + e.Param() + " caracteres.", field)
				case "max":
					invalidParameter = buildInvalidParameterResponse(nil, "O campo " + field + " deve ter no máximo " + e.Param() + " caracteres.", field)
				case "numeric":
					invalidParameter = buildInvalidParameterResponse(nil, "O campo " + field + " deve conter apenas números.", field)
				case "email":
					invalidParameter = buildInvalidParameterResponse(nil, "O campo " + field + " deve ser um e-mail válido.", field)
				case "phone_numeric_format":
					invalidParameter = buildInvalidParameterResponse(nil, "O telefone deve conter apenas números e conter entre 10 e 11 dígitos (DDD + número).", field)
				case "document_validator":
					invalidParameter = buildInvalidParameterResponse(nil, "O documento deve ser válido.", field)
				case "custom_time_format":
					invalidParameter = buildInvalidParameterResponse(nil, "O campo " + field + " deve estar no formato yyyy-MM-dd'T'HH:mm:ss[.SSS]XXX", field)
				case "conflict_date":
					invalidParameter = buildInvalidParameterResponse(nil, "A data hora de termino não pode ser menor ou igual a data hora de inicio .", field)
				default:
					invalidParameter = buildInvalidParameterResponse(nil, "Campo " + field + " inválido.", field)
			}
					
			defaultErrorList = append(defaultErrorList, invalidParameter)
		}
	}
	
	return context_error.ErrorResponse{ErroResponse: defaultErrorList}
}


func buildInvalidParameterResponse(errorMessage *string, description string, parameterName string, ) context_error.DefaultError {
	var invalidParameterErrorMessage = "invalid_parameter"
	
	if errorMessage == nil {
		errorMessage = &invalidParameterErrorMessage
	}
	
	return context_error.DefaultError{
		Error: errorMessage,
		ParameterName: &parameterName,
		Description: &description,
	}
}


func ReadJSON(ctx *gin.Context) ([]byte, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		context_error.JsonReadError(ctx)
		return nil, err
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}


func ValidateJSONAndStruct(ctx *gin.Context, payload any) error {
	if err := ValidateJSON(ctx, payload); err != nil {
		return err
	}

	if err := ValidateStruct(ctx, payload); err != nil {
		return err
	}

	return nil
}


func ValidateJSON(ctx *gin.Context, payload any) error {
	body, err := ReadJSON(ctx)
	if err != nil {
		return err
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := ctx.ShouldBindJSON(payload); err != nil {
		context_error.JsonInvalidError(ctx)
		return err
	}
	return nil
}


func ValidateStruct(ctx *gin.Context, payload any) error {
	if err := Validate.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, TranslateValidationError(err))
		return err
	}

	return nil
}