// models/ScheduleBlock.go
package models

import (
	//"gorm.io/gorm"
)

type ScheduleBlock struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	TemplateID uint           `json:"template_id"`
	Template   ScheduleTemplate `gorm:"foreignKey:TemplateID"`
	StartTime  string          `json:"start_time"` 
	EndTime    string          `json:"end_time"`
	Capacity   int            `json:"capacity"`
	DisciplineID uint         `json:"discipline_id"`
	Discipline Discipline     `gorm:"foreignKey:DisciplineID"`
}
