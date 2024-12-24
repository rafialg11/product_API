package entities

import (
	"time"
)

type Purchase struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	PurchaseDate time.Time `gorm:"not null" json:"purchase_date"`
	TotalAmount  float64   `gorm:"not null" json:"total_amount"`
}
