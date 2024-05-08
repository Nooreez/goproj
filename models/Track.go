package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	track_code string `json:"track_code" gorm:"foreignKey:ID"`
	menu []Menu
}