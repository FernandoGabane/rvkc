package converter

import (
	"time"
	"rvkc/dto"
	"rvkc/models"
	"strings"
	"github.com/google/uuid"
)


func ToConfirmationEntity(dto *dto.ConfirmationRequest, actionAt time.Time) models.Confirmation {

	return models.Confirmation{
		ID:       		strings.ToUpper("CONF_" + uuid.NewString()),
		AccountId: 		*dto.AccountId,
		ClubId:     	*dto.ClubId,
		Status:	 		*dto.Status,
		ActionAt:	    actionAt,
	}
}


func ToConfirmationEntityList(dto []dto.ConfirmationRequest, actionAt time.Time) []models.Confirmation {
	 confirmationEntities := make([]models.Confirmation, len(dto))

	for i, confirmation := range dto {
		confirmationEntities[i] = models.Confirmation{
			ID:       		strings.ToUpper("CONF_" + uuid.NewString()),
			AccountId: 		*confirmation.AccountId,
			ClubId:     	*confirmation.ClubId,
			Status:	 		*confirmation.Status,
			ActionAt:	    actionAt,
		}
	}

	return confirmationEntities
}