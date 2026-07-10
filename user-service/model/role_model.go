package model

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
