package models

import (
	"time"
	//"gorm.io/gorm"
)

type Calendar struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserId          uint		   `json:"user_id"`
	User            User           `json:"user" gorm:"foreignKey:UserId"`
	ClassId         uint           `json:"class_id"`
	Class           Class          `json:"class" gorm:"foreignKey:ClassId"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}