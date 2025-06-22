package services

import (
	"time"
	"wod-go/models"
	"wod-go/database"
	"log"
)

func GenerateClassesFromTemplates(gymID uint, fromDate, toDate time.Time) error {
	var templates []models.ScheduleTemplate

	// Traer templates del gimnasio
	if err := database.DB.
		Preload("Blocks").
		Where("gym_id = ?", gymID).
		Find(&templates).Error; err != nil {
		return err
	}

	// Mapa de días en inglés a español
	dayMap := map[string]string{
		"Monday":    "Lunes",
		"Tuesday":   "Martes",
		"Wednesday": "Miércoles",
		"Thursday":  "Jueves",
		"Friday":    "Viernes",
		"Saturday":  "Sábado",
		"Sunday":    "Domingo",
	}

	// Recorrer los días desde from hasta to
	for d := fromDate; !d.After(toDate); d = d.AddDate(0, 0, 1) {
		weekday := dayMap[d.Weekday().String()] // Convertir al día en español

		// Buscar templates que coincidan con ese día
		for _, template := range templates {
			if template.Day == weekday {
				// Por cada bloque de ese template, crear una clase
				for _, block := range template.Blocks {
					class := models.Class{
						Date:         d.Format("2006-01-02"),
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
