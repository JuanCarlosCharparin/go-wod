package models

import (
	"time"
	"gorm.io/gorm"
)

type UserPack struct {
	Id              uint           `json:"id" gorm:"primaryKey"`
	StartDate       time.Time      `json:"start_date"`
	ExpirationDate  time.Time      `json:"expiration_date"`
	Status          int            `json:"status"`
	GymId           uint           `json:"gym_id"`           
	Gym             Gym            `json:"gym" gorm:"foreignKey:GymId"`
	UserId          uint           `json:"user_id"`             
	User            User           `json:"user" gorm:"foreignKey:UserId"`
	PackId          uint           `json:"pack_id"`             
	Pack            Pack           `json:"pack" gorm:"foreignKey:PackId"`
	DisciplineId    uint           `json:"discipline_id"`             
	Discipline      Discipline     `json:"discipline" gorm:"foreignKey:DisciplineId"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (UserPack) TableName() string {
	return "users_packs"
}