package models

type Product struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Category    string      `json:"category"`
	Inventory   []Inventory `gorm:"foreignKey:ProductID" json:"inventory,omitempty"`
	Orders      []Order     `gorm:"foreignKey:ProductID" json:"orders,omitempty"`
}

type Inventory struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Location  string  `json:"location"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}

type Order struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	OrderDate string  `json:"order_date"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}
