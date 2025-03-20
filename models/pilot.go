package models

import (
	"strings"
)

type Pilot struct {
    ID        uint   `gorm:"primaryKey" json:"-"`
    Document  string `gorm:"unique;not null" json:"document"`
    Name      string `gorm:"not null" json:"name"`
    Phone     string `gorm:"not null" json:"phone"`
    Email     string `gorm:"not null" json:"email"`
}

func (Pilot) TableName() string {
    return "pilot"
}

func (p *Pilot) Higienize() {
	p.Name     = strings.ToUpper(strings.Join(strings.Fields(p.Name), " "))
	p.Document = strings.TrimSpace(p.Document)
	p.Phone    = strings.TrimSpace(p.Phone)
	p.Email    = strings.TrimSpace(strings.ToLower(p.Email))
}
