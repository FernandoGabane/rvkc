package dto

import (
	"rvkc/middleware"
	"rvkc/models"
)


type ConfirmationResponse struct {
	ID	   	  *string 		   `json:"id"`
	AccountId *string 		   `json:"account_id,omitempty"`
	ClubId    *string 		   `json:"club_id,omitempty"`
	Status 	  *string 		   `json:"status"`
	Account   *AccountResponse `json:"account,omitempty"`
	Club      *ClubResponse    `json:"club,omitempty"`
}


func ToConfirmationResponse(confirmation *models.Confirmation) ConfirmationResponse {
	var accountResponse *AccountResponse
	var clubResponse *ClubResponse
	
	if confirmation.Account.ID != "" {
		accountResponse = &AccountResponse{
			ID:       confirmation.Account.ID,
			Name:     confirmation.Account.Name,
		}
	}

	if confirmation.Club.ID != "" {
		clubResponse = &ClubResponse{
			ID:        &confirmation.Club.ID,
			Name:      &confirmation.Club.Name,
			StartAt:   &middleware.CustomTime{ Time: confirmation.Club.StartAt, ParseError: nil},
			EndAt:     &middleware.CustomTime{ Time: confirmation.Club.EndAt, ParseError: nil},
			AccountId: &confirmation.Club.AccountId,
			Weekday:   &confirmation.Club.Weekday,
			Slots:     &confirmation.Club.Slots,
		}
	} 
	
	return ConfirmationResponse{
		ID:        &confirmation.ID,
		AccountId: &confirmation.AccountId,
		ClubId:    &confirmation.ClubId,
		Status:    &confirmation.Status,
		Account:   accountResponse,
		Club:      clubResponse,
	}
}


func ToConfirmationResponseList(confirmations []*models.Confirmation) []ConfirmationResponse {
	var confirmationResponseList []ConfirmationResponse = make([]ConfirmationResponse, 0)	

	for _, confirmation 	 	 := range confirmations {
		confirmationResponse 	 := ToConfirmationResponse(confirmation)
		confirmationResponseList = append(confirmationResponseList, confirmationResponse)
	}

	return confirmationResponseList
}
