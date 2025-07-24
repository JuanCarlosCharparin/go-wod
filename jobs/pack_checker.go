package jobs

import (
	"log"
	"time"
	"wod-go/database"
	"wod-go/models"
	"wod-go/services"
)

//go run main.go check-expired-packs

func CheckExpiredUserPacks() {
	var userPacks []models.UserPack

	// Traer todos los packs activos
	err := database.DB.Preload("Pack").
		Where("status = ?", 1).
		Find(&userPacks).Error

	if err != nil {
		log.Println("Error obteniendo packs:", err)
		return
	}

	for _, pack := range userPacks {
		now := time.Now()

		// Condición 1: Fecha vencida
		expired := now.After(pack.ExpirationDate)

		// Condición 2: Clases agotadas
		used, err := services.CountUsedClasses(pack.UserId, pack.GymId, pack.DisciplineId, pack.StartDate, pack.ExpirationDate)
		if err != nil {
			log.Printf("Error contando clases para pack_id=%d: %v\n", pack.Id, err)
			continue
		}

		usedAll := used >= pack.Pack.ClassQuantity

		if expired || usedAll {
			pack.Status = 0
			database.DB.Save(&pack)
			log.Printf("Pack vencido para user_id=%d (pack_id=%d)\n", pack.UserId, pack.Id)
		}
	}
}