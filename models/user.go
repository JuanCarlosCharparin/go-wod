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
	Password        string         `json:"-"`
	GymId           *uint           `json:"gym_id"`
	Gym             Gym            `json:"gym" gorm:"foreignKey:GymId"`
	RoleId 			uint           `json:"role_id"` 
	Role   			Role           `json:"role" gorm:"foreignKey:RoleId"`
	Status          bool           `json:"status" gorm:"default:true"`
	UserPacks      []UserPack     `gorm:"foreignKey:UserId"` //relacion con UserPack
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}