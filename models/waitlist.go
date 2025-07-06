package models

import (
	"time"
	"gorm.io/gorm"
)


type Waitlist struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	ClassId   uint      `json:"class_id"`
	Class     Class     `gorm:"foreignKey:ClassId"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
