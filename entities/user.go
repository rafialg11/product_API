package entities

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
