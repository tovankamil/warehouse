package model

import "time"

type Role struct {
	ID         uint      `json:"id" gorm:"primaryKey:autoincrement"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
