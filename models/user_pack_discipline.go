package models

import (
	"time"
	"gorm.io/gorm"
)

type UserPackDiscipline struct {
	Id              uint           `json:"id" gorm:"primaryKey"`
	UserPackId      uint           `json:"user_pack_id"`             
	UserPack        UserPack       `json:"user_pack" gorm:"foreignKey:UserPackId"`
	DisciplineId    uint           `json:"discipline_id"`             
	Discipline      Discipline     `json:"discipline" gorm:"foreignKey:DisciplineId"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}