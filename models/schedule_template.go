package models

import (
	//"gorm.io/gorm"
)

type ScheduleTemplate struct {
	ID      uint           `gorm:"primaryKey" json:"id"`
	GymID   uint           `json:"gym_id"`
	Gym     Gym            `gorm:"foreignKey:GymID"`
	Day     string         `json:"day"`
	Blocks  []ScheduleBlock `gorm:"foreignKey:TemplateID" json:"blocks"`
}