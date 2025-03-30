package converter

import (
	"rvkc/dto"
	"rvkc/models"
)

func ToAccountEntity(dto *dto.AccountRequest) models.Account {
	return models.Account{
		Document: *dto.Document,
		Name:     *dto.Name,
		Phone:    *dto.Phone,
		Email:    *dto.Email,
	}
}