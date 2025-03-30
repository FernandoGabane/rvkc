package converter

import (
	"log"
	"time"

	"rvkc/dto"
	"rvkc/models"
)

func ToClubEntity(clubRequest *dto.ClubRequest) models.Club {
	dateLayout := "2006-01-02"
	timeLayout := "15:04"

	parsedDate, err := time.Parse(dateLayout, *clubRequest.Date)
	if err != nil {
		log.Printf("Erro ao fazer parse da data: %v", err)
	}

	startAt, err := time.Parse(timeLayout, *clubRequest.StartAt)
	if err != nil {
		log.Printf("Erro ao fazer parse do horário de início: %v", err)
	}

	endAt, err := time.Parse(timeLayout, *clubRequest.EndAt)
	if err != nil {
		log.Printf("Erro ao fazer parse do horário de término: %v", err)
	}

	return models.Club{
		Name:    *clubRequest.Name,
		Date:    parsedDate.Format("2006-01-02"),
		Weekday: translateWeekday(parsedDate),
		StartAt: startAt.Format("15:04:05"),
		EndAt:   endAt.Format("15:04:05"),
		Slots:   *clubRequest.Slots,
	}
}

func translateWeekday(date time.Time) string {
	switch date.Weekday() {
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
	default:
		return "DOMINGO"
	}
}
