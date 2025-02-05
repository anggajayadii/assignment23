package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Location  string `json:"location"`
}
