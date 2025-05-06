package models

import (
	"strings"
	"time"
)

type Club struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null"   json:"name"`
	Weekday   string    `gorm:"not null"   json:"weekday"`
	StartAt   time.Time `gorm:"not null"   json:"start_at"`
	EndAt     time.Time `gorm:"not null"   json:"end_at"`
	AccountId string  	`gorm:"not null"   json:"account_id"`
	Slots     uint   	`gorm:"not null"   json:"slots"`
}

func (Club) TableName() string {
	return "club"
}

func (p *Club) Higienize() {
	p.Name = strings.ToUpper(strings.Join(strings.Fields(p.Name), " "))
}