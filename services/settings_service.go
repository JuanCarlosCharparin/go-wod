package services

import (
	"wod-go/database"
	"wod-go/models"
	"errors"
)

func GetGymCancelTimeLimit(gymID uint) (int, error) {
	var setting models.GymSetting
	if err := database.DB.
		Where("gym_id = ?", gymID).
		First(&setting).Error; err != nil {
		return 0, errors.New("no se encontro configuracion para el gimnasio")
	}
	return setting.CancelTimeLimitMinutes, nil
}
