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

	// Usamos la fecha actual, no la fecha de la clase
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

	// Contar clases ya tomadas dentro del período del pack
	var used int64
	database.DB.Model(&models.Calendar{}).
		Where("user_id = ? AND class_id IN (?)",
			userID,
			database.DB.Model(&models.Class{}).
				Select("id").
				Where("date BETWEEN ? AND ? AND gym_id = ? AND discipline_id = ?",
					userPack.StartDate.Format("2006-01-02"), userPack.ExpirationDate.Format("2006-01-02"), gymID, disciplineID),
		).Count(&used)

	return &PackUsage{
		PackID:        userPack.PackId,
		ClassQuantity: userPack.Pack.ClassQuantity,
		Used:          int(used),
		Remaining:     userPack.Pack.ClassQuantity - int(used),
	}, nil
}