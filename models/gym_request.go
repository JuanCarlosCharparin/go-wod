package models

import (
	"time"
	"gorm.io/gorm"
)

type GymRequest struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint
	User      User      `gorm:"foreignKey:UserId"`
	GymId     uint
	Gym       Gym       `gorm:"foreignKey:GymId"`
	Status    string    
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}