package models

import (
	"strings"
)

type Club struct {
    ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"not null"   json:"name"`
    Recurrence string `gorm:"not null"   json:"recurrence"`
	Weekday    string `gorm:"not null"   json:"weekday"`
	StartAt    string `gorm:"not null"   json:"start_at"`
	EndAt      string `gorm:"not null"   json:"end_at"`
}

func (Club) TableName() string {
    return "club"
}

func (p *Club) Higienize() {
	p.Name       = strings.ToUpper(strings.Join(strings.Fields(p.Name), " "))
	p.Recurrence = strings.TrimSpace(p.Recurrence)
	p.Weekday    = strings.TrimSpace(p.Weekday)
	p.StartAt    = strings.TrimSpace(p.StartAt)
	p.EndAt      = strings.TrimSpace(p.EndAt)
}
