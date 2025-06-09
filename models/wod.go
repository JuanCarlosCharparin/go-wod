package models

import (
	"time"
	"gorm.io/gorm"
)

type Wod struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Type            string         `json:"type"`
	Duration        int            `json:"duration"`
	Level           string         `json:"level"`
	GymId           uint           `json:"gym_id"`
	Gym             Gym            `json:"gym" gorm:"foreignKey:GymId"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}