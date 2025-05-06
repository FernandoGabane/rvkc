package dto

import (
	"rvkc/middleware"
	"rvkc/models"
)

type ClubResponse struct {
	ID        *string     		   	 `json:"id,omitempty"`
	Name      *string     		   	 `json:"name,omitempty"`
	StartAt   *middleware.CustomTime `json:"start_at,omitempty"`
	EndAt  	  *middleware.CustomTime `json:"end_at,omitempty"`
	Weekday   *string     		   	 `json:"weekday,omitempty"`
	AccountId *string     		   	 `json:"account_id,omitempty"`
	Slots     *uint       		   	 `json:"slots,omitempty"`
}


func ToClubResponse(club *models.Club) ClubResponse {

	startAt := middleware.CustomTime{
		Time: club.StartAt,
	}

	endAt := middleware.CustomTime{
		Time: club.EndAt,
	}

	return ClubResponse{
		ID: 	   &club.ID,
		Name:      &club.Name,
		StartAt:   &startAt,
		EndAt:     &endAt,
		AccountId: &club.AccountId,
		Weekday:   &club.Weekday,
		Slots:     &club.Slots,
	}
}


func ToClubResponseList(clubs []*models.Club) []ClubResponse {
	var clubResponseList []ClubResponse = make([]ClubResponse, 0)
	
	for _, club := range clubs {
		clubResponse 	:= ToClubResponse(club)
		clubResponseList = append(clubResponseList, clubResponse)
	}


	return clubResponseList
}
