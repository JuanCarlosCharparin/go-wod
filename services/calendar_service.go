package services

import (
	"log"
	"time"
	"wod-go/database"
	"wod-go/models"
)

type PackUsage struct {
	PackID        uint
	ClassQuantity int
	Used          int
	Remaining     int
}

func GetPackUsage(userID, gymID uint, disciplineIDs []uint, classDate time.Time) (*PackUsage, error) {
	var userPack models.UserPack
	today := time.Now()

	// Buscar el pack activo del usuario para el gimnasio y disciplinas
	err := database.DB.Preload("Pack").Preload("Disciplines").
		Where("user_id = ? AND gym_id = ? AND start_date <= ? AND expiration_date >= ? AND status = 1",
			userID, gymID, today, today).
		First(&userPack).Error

	if err != nil {
		log.Printf("No se encontró un pack válido para user_id=%d, gym_id=%d y fecha actual=%s\n",
			userID, gymID, today.Format("2006-01-02"))
		return nil, err
	}

	used, err := CountUsedClasses(userID, gymID, disciplineIDs, userPack.StartDate, userPack.ExpirationDate)
	if err != nil {
		log.Printf("Error al contar clases usadas: %v", err)
		return nil, err
	}

	// Determinar la cantidad de clases según el caso
    var classQuantity int
    var packID uint
    if userPack.PackId != nil && userPack.Pack != nil {
        classQuantity = userPack.Pack.ClassQuantity
        packID = *userPack.PackId
    } else if userPack.ClassQuantity != nil {
        classQuantity = *userPack.ClassQuantity
        packID = 0
    } else {
        // No hay pack ni cantidad de clases definida
        log.Printf("El pack no tiene cantidad de clases definida para user_id=%d, gym_id=%d", userID, gymID)
    	return nil, nil
    }
	return &PackUsage{
		PackID:        packID,
		ClassQuantity: classQuantity,
		Used:          used,
		Remaining:     classQuantity - used,
	}, nil
}

// Contar clases de un usuario para múltiples disciplinas
func CountUsedClasses(userID, gymID uint, disciplineIDs []uint, startDate, endDate time.Time) (int, error) {
	var used int64
	err := database.DB.Model(&models.Calendar{}).
		Where("user_id = ? AND class_id IN (?) AND status IN ?", 
			userID,
			database.DB.Model(&models.Class{}).
				Select("id").
				Where("date BETWEEN ? AND ? AND gym_id = ? AND discipline_id IN ?", 
					startDate, endDate, gymID, disciplineIDs),
			[]string{"inscripto", "ausente"},
		).
		Distinct("class_id").
		Count(&used).Error

	return int(used), err
}

