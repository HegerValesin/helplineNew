package models

import (
	"time"
)

type UserType string

const (
	Admin   UserType = "admin"
	Teacher UserType = "teacher"
)

type User struct {
	ID         string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Email      string    `json:"email"`
	Type       UserType  `json:"type"`
	Password   string    `json:"-"`
	Facilities []string  `json:"facilities" gorm:"type:text[]"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	CreatedBy  string    `json:"createdBy"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	UpdatedBy  string    `json:"updatedBy"`
}
