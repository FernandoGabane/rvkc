package models

import (
	"strings"
)

type Account struct {
    ID          string          `gorm:"primaryKey"              json:"id"`
    Document    string          `gorm:"unique       ;not null"  json:"document"`
    Name        string          `gorm:"not null"                json:"name"`
    Phone       string          `gorm:"not null"                json:"phone"`
    Email       string          `gorm:"not null"                json:"email"`
    Roles       []*Role         `gorm:"many2many:account_role;" json:"roles"`
}

func (Account) TableName() string {
    return "account"
}

func (p *Account) Higienize() {
	p.Name     = strings.ToUpper(strings.Join(strings.Fields(p.Name), " "))
	p.Document = strings.TrimSpace(p.Document)
	p.Phone    = strings.TrimSpace(p.Phone)
	p.Email    = strings.TrimSpace(strings.ToLower(p.Email))
}
