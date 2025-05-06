package converter

import (
	"strings"
	"time"
	"rvkc/dto"
	"rvkc/models"
	
	"github.com/google/uuid"
)


func ToClubEntity(clubRequest *dto.ClubRequest) models.Club {
	return models.Club{
		ID: 		strings.ToUpper("CLUB_" + uuid.NewString()),
		Name:   	*clubRequest.Name,
		Weekday: 	translateWeekday(clubRequest.StartAt.Time),
		StartAt: 	clubRequest.StartAt.Time,
		EndAt:   	clubRequest.EndAt.Time,
		AccountId: 	*clubRequest.AccountId,
		Slots:   	*clubRequest.Slots,
	}
}


func translateWeekday(t time.Time) string {
	switch t.Weekday() {
	case time.Monday:
		return "SEGUNDA-FEIRA"
	case time.Tuesday:
		return "TERÇA-FEIRA"
	case time.Wednesday:
		return "QUARTA-FEIRA"
	case time.Thursday:
		return "QUINTA-FEIRA"
	case time.Friday:
		return "SEXTA-FEIRA"
	case time.Saturday:
		return "SÁBADO"
	case time.Sunday:
		return "DOMINGO"
	default:
		return ""
	}
}
