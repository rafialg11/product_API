package entities

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"size:255; not null" json:"name"`
	Description string    `gorm:"type:text; not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	CreatedAt   time.Time `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
