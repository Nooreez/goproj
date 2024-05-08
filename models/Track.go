package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	TrackCode string `json:"track_code"`
	Menu      []Menu `json:"menus"`
}
