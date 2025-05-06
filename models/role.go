package models

type Role struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

func (Role) TableName() string {
	return "role"
}
