package models

type Image struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FileName  string `json:"file_name"`
	ProductID uint   `json:"product_id"`
}
