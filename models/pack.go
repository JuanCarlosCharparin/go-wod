package models

import (
	"time"
	"gorm.io/gorm"
)

type Pack struct {
	Id             uint           `json:"id" gorm:"primaryKey"`
	PackName       string         `json:"pack_name"`
	Price          float64        `json:"price"`              
	ClassQuantity  int            `json:"class_quantity"`     
	Months         int            `json:"months"`             
	GymId          uint           `json:"gym_id"`             
	Gym            Gym            `json:"gym" gorm:"foreignKey:GymId"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Pack) TableName() string {
	return "packs"
}