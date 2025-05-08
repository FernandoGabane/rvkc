package models

import (
	"time"
)

type Confirmation struct {
	ID       		string    `gorm:"primaryKey" 		   json:"id"`
	AccountId   	string    `gorm:"not null"  		   json:"account_id"`
	ClubId 		   	string    `gorm:"not null"   		   json:"club_id"`
	Status          string    `gorm:"not null"   		   json:"action"`
	ActionAt 		time.Time `gorm:"not null"   		   json:"action_at"`
	Account 		Account   `gorm:"foreignKey:AccountId" json:"account"`
	Club    		Club      `gorm:"foreignKey:ClubId"    json:"club"`
}

func (Confirmation) TableName() string {
	return "confirmation"
}