package models

import (
	"time"
)

type OccurrenceStatus string

const (
	Open       OccurrenceStatus = "open"
	InProgress OccurrenceStatus = "in_progress"
	Closed     OccurrenceStatus = "closed"
)

type Occurrence struct {
	ID          string            `json:"id" gorm:"primaryKey"`
	Date        time.Time         `json:"date"`
	Description string            `json:"description"`
	Status      OccurrenceStatus  `json:"status"`
	AssignedTo  string            `json:"assignedTo"`
	Observation string            `json:"observation"`
	ServiceID   string            `json:"serviceId"`
	UserID      string            `json:"userId"`
	HasAudio    bool              `json:"hasAudio"`
	FacilityID  string            `json:"facilityId"`
	FloorID     string            `json:"floorId"`
	RoomID      string            `json:"roomId"`
	Audios      []OccurrenceAudio `json:"audios" gorm:"foreignKey:OccurrenceID"`
	CreatedAt   time.Time         `json:"createdAt"`
	CreatedBy   string            `json:"createdBy"`
	UpdatedAt   time.Time         `json:"updatedAt"`
	UpdatedBy   string            `json:"updatedBy"`
}
