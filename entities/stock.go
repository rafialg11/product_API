package entities

type Stock struct {
	ID          uint `gorm:"primary_key" json:"id"`
	ProductID   uint `gorm:"not null" json:"product_id"`
	Quantity    uint `gorm:"not null" json:"quantity"`
	PurchaseQty uint `gorm:"not null" json:"purchase_qty"`
	SaleQty     uint `gorm:"not null" json:"sale_qty"`
}
