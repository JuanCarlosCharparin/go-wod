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

func GetPackUsage(userID, gymID, disciplineID uint, classDate time.Time) (*PackUsage, error) {
	var userPack models.UserPack
	today := time.Now()

	err := database.DB.Preload("Pack").
		Where("user_id = ? AND gym_id = ? AND discipline_id = ? AND start_date <= ? AND expiration_date >= ? AND status = 1",
			userID, gymID, disciplineID, today, today).
		First(&userPack).Error

	if err != nil {
		log.Printf("No se encontró un pack válido para user_id=%d, gym_id=%d, disciplina=%d y fecha actual=%s\n",
			userID, gymID, disciplineID, today.Format("2006-01-02"))
		return nil, err
	}

	used, err := CountUsedClasses(userID, gymID, disciplineID, userPack.StartDate, userPack.ExpirationDate)
	if err != nil {
		log.Printf("Error al contar clases usadas: %v", err)
		return nil, err
	}

	return &PackUsage{
		PackID:        userPack.PackId,
		ClassQuantity: userPack.Pack.ClassQuantity,
		Used:          used,
		Remaining:     userPack.Pack.ClassQuantity - used,
	}, nil
}


// Contar clases de un usuario
func CountUsedClasses(userID, gymID, disciplineID uint, startDate, endDate time.Time) (int, error) {
	var used int64
	err := database.DB.Model(&models.Calendar{}).
		Where("user_id = ? AND class_id IN (?) and status = 'inscripto'",
			userID,
			database.DB.Model(&models.Class{}).
				Select("id").
				Where("date BETWEEN ? AND ? AND gym_id = ? AND discipline_id = ?",
					startDate, endDate, gymID, disciplineID),
		).Count(&used).Error
	return int(used), err
}