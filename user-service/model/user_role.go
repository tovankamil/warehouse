package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;ForeignKey:UserID;references:ID"`
	RoleID     uuid.UUID `json:"role_id" gorm:"type:uuid;ForeignKey:RoleID;references:ID"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Table interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "user_role"
}
