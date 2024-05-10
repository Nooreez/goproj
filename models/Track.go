package models

type Track struct {
	TrackCode string `gorm:"unique" json:"track_code"`
	MenuItem  string `json:"menu_item"`
}
