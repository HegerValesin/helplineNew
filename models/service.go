package models

import (
	"time"
)

type ServiceStatus string

const (
	Active   ServiceStatus = "ativo"
	Inactive ServiceStatus = "inativo"
)

type Service struct {
	ID          string        `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      ServiceStatus `json:"status"`
	SectorID    string        `json:"sectorId"`
	CreatedAt   time.Time     `json:"createdAt"`
	CreatedBy   string        `json:"createdBy"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	UpdatedBy   string        `json:"updatedBy"`
}
