package models

import (
	"time"
	"gorm.io/gorm"
)

type Discipline struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	GymId     uint           `json:"gym_id"`
	Gym       Gym            `json:"gym" gorm:"foreignKey:GymId"` //relacion a gym_id
}