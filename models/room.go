package models

import (
	"time"
)

type Room struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	RoomNumber int       `json:"roomNumber"`
	Area       float64   `json:"area"`
	FloorID    string    `json:"floorId"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UpdatedBy  string    `json:"updatedBy"`
}
