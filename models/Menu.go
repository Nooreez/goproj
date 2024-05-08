package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	ItemName string  `json:"item_name"`
	Price    float64 `json:"price"`
}
