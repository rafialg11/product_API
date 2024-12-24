package entities

type PurchaseDetail struct {
	ID           uint    `gorm:"primary_key" json:"id"`
	PurchaseID   uint    `gorm:"not null" json:"purchase_id"`
	ProductID    uint    `gorm:"not null" json:"product_id"`
	Quantity     uint    `gorm:"not null" json:"quantity"`
	PricePerUnit float64 `gorm:"not null" json:"price_per_unit"`
}
