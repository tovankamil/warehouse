package model

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey:autoincrement"`
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
