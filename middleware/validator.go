package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"rvkc/dto"
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

func TranslateValidationError(err error) map[string]string {
	errorsMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(e.Field(), "${1}_${2}")
			field = strings.ToLower(field)

			switch e.Tag() {
			case "required":
				errorsMap[field] = "O campo " + field + " é obrigatório."
			case "len":
				errorsMap[field] = "O campo " + field + " deve ter exatamente " + e.Param() + " caracteres."
			case "min":
				errorsMap[field] = "O campo " + field + " deve ter no mínimo " + e.Param() + " caracteres."
			case "max":
				errorsMap[field] = "O campo " + field + " deve ter no máximo " + e.Param() + " caracteres."
			case "numeric":
				errorsMap[field] = "O campo " + field + " deve conter apenas números."
			case "email":
				errorsMap[field] = "O campo " + field + " deve ser um e-mail válido."
			case "telefone_numeric":
				errorsMap[field] = "O telefone deve conter apenas números e ter entre 8 e 15 dígitos."
			case "time_format":
				errorsMap[field] = "O campo " + field + " deve estar no formato HH:MM."
			case "conflict_date":
				errorsMap[field] = "O end_at não pode ser menor ou igual ao start_at."
			default:
				errorsMap[field] = "Campo " + field + " inválido."
			}
		}
	}

	return errorsMap
}

func ReadJSON(ctx *gin.Context) ([]byte, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler o JSON"})
		return nil, err
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

func ValidateJSON(ctx *gin.Context, payload any) error {
	body, err := ReadJSON(ctx)
	if err != nil {
		return err
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return err
	}
	return nil
}

func ValidateStruct(ctx *gin.Context, payload any) error {
	if err := dto.Validate.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": TranslateValidationError(err)})
		return err
	}

	return nil
}