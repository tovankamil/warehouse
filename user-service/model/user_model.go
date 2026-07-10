package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Photo      string    `json:"photo"`
	Phone      string    `json:"phone"`
	Roles      []Role    `gorm:"many2many:user_role"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time
}
