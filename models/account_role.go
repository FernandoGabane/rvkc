package models

type AccountRole struct {
	AccountID string        `gorm:"primaryKey" json:"account_id"`
	RoleID    string        `gorm:"primaryKey" json:"role_id"`
	Account   Account       `gorm:"foreignKey: AccountID;references:ID"`
	Role      Role          `gorm:"foreignKey: RoleID;references:ID"`
}

func (AccountRole) TableName() string {
	return "account_role"
}
