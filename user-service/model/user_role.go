package model

import "time"

type UserRole struct {
	ID         uint      `json:"id" gorm:"primaryKey:autoincrement"`
	UserID     uint      `json:"user_id" gorm:"ForeignKey:UserID;references:ID"`
	RoleID     uint      `json:"role_id" gorm:"ForeignKey:RoleID;references:ID"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Table interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "user_role"
}
