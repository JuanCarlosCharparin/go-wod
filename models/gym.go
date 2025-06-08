package models

import (
	"time"
	"gorm.io/gorm"
)

type Gym struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Location  string         `json:"location"`
	Phone  	  string         `json:"phone"`
	Email  	  string         `json:"email"`
	CountryId  uint           `json:"country_id"`                          // Clave foránea
	Country    Country        `json:"country" gorm:"foreignKey:CountryId"` // Relación
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}