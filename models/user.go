package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	Id              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name"`
	Lastname        string         `json:"lastname"`
	Gender          string         `json:"gender"`
	Phone           string         `json:"phone"`
	Email           string         `json:"email" gorm:"unique"`
	MovilPhone      *string        `json:"movil_phone"`
	BirthDate       *string        `json:"birth_date"`
	DNI             *string        `json:"dni"`
	Password        string         `json:"password"`
	GymId           uint           `json:"gym_id"`
	Gym             Gym            `json:"gym" gorm:"foreignKey:GymId"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}