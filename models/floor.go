package models

import (
	"time"
)

type Floor struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	FloorNumber int       `json:"floorNumber"`
	FacilityID  string    `json:"facilityId"`
	Rooms       []string  `json:"rooms" gorm:"type:text[]"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedBy   string    `json:"updatedBy"`
}
