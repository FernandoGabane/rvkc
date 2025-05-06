package dto


type ConfirmationRequest struct {
	AccountId      *string     			  `json:"account_id"      validate:"required,max=100"`
	ClubId    	   *string     			  `json:"club_id"         validate:"required,max=100"`
	Status 		   *string                `json:"status"          validate:"required,oneof=CONFIRMED UNCONFIRMED PENDING"`
}


type ConfirmationRequestList struct {
	Confirmations []ConfirmationRequest `json:"confirmations" validate:"required,dive"`
}

const (
	CONFIRMED  = "CONFIRMED"
	CANCELLED  = "UNCONFIRMED"
	PENDING    = "PENDING"
)