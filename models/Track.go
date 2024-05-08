package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	TrackCode string `json:"track_code"`
	Menu      []Menu `gorm:"many2many:track_menu_items;" json:"menu"`
}
