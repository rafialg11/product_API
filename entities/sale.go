package entities

import (
	"time"
)

type Sale struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	SaleDate    time.Time `gorm:"not null" json:"sale_date"`
	TotalAmount float64   `gorm:"not null" json:"total_amount"`
}
