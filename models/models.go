package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Category    string      `json:"category"`
	Inventory   []Inventory `json:"inventory"` // Relasi one-to-many dengan Inventory
	Orders      []Order     `json:"orders"`    // Relasi one-to-many dengan Order
}
type Inventory struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Location  string  `json:"location"`
	Product   Product `gorm:"foreignkey:ProductID;association_foreignkey:ID"` // Relasi belongs-to dengan Product
}

type Order struct {
	gorm.Model
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	OrderDate string  `json:"order_date"`
	Product   Product `gorm:"foreignkey:ProductID;association_foreignkey:ID"` // Relasi belongs-to dengan Product
}
