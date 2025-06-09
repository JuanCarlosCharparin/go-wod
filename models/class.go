package models

import (
	"time"
	"gorm.io/gorm"
)

type Class struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Date            time.Time      `json:"date"`
	Time            string         `json:"time"`
	Capacity        int            `json:"capacity"`
	GymId           uint           `json:"gym_id"`
	Gym             Gym            `json:"gym" gorm:"foreignKey:GymId"`
	DisciplineId    uint           `json:"discipline_id"`
	Discipline      Discipline     `json:"discipline" gorm:"foreignKey:DisciplineId"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}