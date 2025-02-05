package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	OrderDate string `json:"order_date"`
}
