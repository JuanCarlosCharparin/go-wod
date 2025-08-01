package models

import (
	"time"
	//"gorm.io/gorm"
)

type Calendar struct {
	Id              uint           `json:"id" gorm:"primaryKey"`
	UserId          uint		   `json:"user_id"`
	User            User           `json:"user" gorm:"foreignKey:UserId"`
	ClassId         uint           `json:"class_id"`
	Class           Class          `json:"class" gorm:"foreignKey:ClassId"`
	Status          string		   `json:"status"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (Calendar) TableName() string {
	return "calendar"
}