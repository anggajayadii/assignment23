package models

type Order struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	OrderDate string  `json:"order_date"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}
