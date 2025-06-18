package models

import "time"

type Student struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	NIM       string    `json:"nim" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Major     string    `json:"major"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
