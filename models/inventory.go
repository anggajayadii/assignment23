package models

type Inventory struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Location  string  `json:"location"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}
