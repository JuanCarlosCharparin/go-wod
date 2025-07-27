package services

import (
	"log"
	"time"
	"wod-go/database"
	"wod-go/models"

	"gorm.io/gorm"
)

func GenerateClassesFromTemplates(gymID uint, fromDate, toDate time.Time) error {
	var templates []models.ScheduleTemplate

	if err := database.DB.
		Preload("Blocks").
		Where("gym_id = ?", gymID).
		Find(&templates).Error; err != nil {
		return err
	}

	dayMap := map[string]string{
		"Monday":    "Lunes",
		"Tuesday":   "Martes",
		"Wednesday": "Miércoles",
		"Thursday":  "Jueves",
		"Friday":    "Viernes",
		"Saturday":  "Sábado",
		"Sunday":    "Domingo",
	}

	for d := fromDate; !d.After(toDate); d = d.AddDate(0, 0, 1) {
		weekday := dayMap[d.Weekday().String()]

		for _, template := range templates {
			if template.Day == weekday {
				for _, block := range template.Blocks {
					dateStr := d.Format("2006-01-02")

					// Verificar si ya existe una clase con ese horario
					var existing models.Class
					err := database.DB.
						Where("date = ? AND time = ? AND gym_id = ? AND discipline_id = ?", dateStr, block.StartTime, gymID, block.DisciplineID).
						First(&existing).Error

					if err == nil {
						// Ya existe, no la creamos
						continue
					} else if err != gorm.ErrRecordNotFound {
						// Error inesperado
						log.Printf("Error buscando clase existente: %v\n", err)
						continue
					}

					// No existe, la creamos
					class := models.Class{
						Date:         dateStr,
						Time:         block.StartTime,
						Capacity:     block.Capacity,
						GymId:        gymID,
						DisciplineId: block.DisciplineID,
					}
					if err := database.DB.Create(&class).Error; err != nil {
						log.Printf("Error creando clase: %v\n", err)
					}
				}
			}
		}
	}

	return nil
}
