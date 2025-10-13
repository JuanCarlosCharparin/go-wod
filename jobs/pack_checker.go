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
	err := database.DB.Preload("Pack").Preload("Disciplines.Discipline").
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

		// Extraer DisciplineIDs
		var disciplineIDs []uint
		for _, upd := range pack.Disciplines {
			disciplineIDs = append(disciplineIDs, upd.DisciplineId)
		}


		// Determinar la cantidad de clases según el caso
		var classQuantity int
		if pack.PackId != nil && pack.Pack != nil {
			classQuantity = pack.Pack.ClassQuantity
		} else if pack.ClassQuantity != nil {
			classQuantity = *pack.ClassQuantity
		} else {
			log.Printf("Pack sin cantidad de clases definida para user_id=%d (pack_id=%d)\n", pack.UserId, pack.Id)
			continue
		}

		// Condición 2: Clases agotadas
		used, err := services.CountUsedClasses(pack.UserId, pack.GymId, disciplineIDs, pack.StartDate, pack.ExpirationDate)
		if err != nil {
			log.Printf("Error contando clases para pack_id=%d: %v\n", pack.Id, err)
			continue
		}

		usedAll := used >= classQuantity

		if expired || usedAll {
			pack.Status = 0
			database.DB.Save(&pack)
			log.Printf("Pack vencido para user_id=%d (pack_id=%d)\n", pack.UserId, pack.Id)
		}
	}
}