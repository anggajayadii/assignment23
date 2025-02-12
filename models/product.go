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
