package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Status      string `json:"status"`
	TrackCode   string `gorm:"unique" json:"track_code"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}
