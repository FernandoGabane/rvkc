package converter

import (
	"rvkc/dto"
	"rvkc/models"
	"strings"
	"github.com/google/uuid"
)

func ToAccountEntity(dto *dto.AccountRequest) models.Account {
	if dto.Roles == nil {
		dto.Roles = []*models.Role{}
	}

	roles := make([]*models.Role, 0, len(dto.Roles))
	for _, role := range dto.Roles {
		roles = append(roles, &models.Role{
			Name: role.Name,
		})
	}

	return models.Account{
		ID:       strings.ToUpper("ACCO_" + uuid.NewString()),
		Document: *dto.Document,
		Name:     *dto.Name,
		Phone:    *dto.Phone,
		Email:    *dto.Email,
		Roles:    roles,
	}
}
