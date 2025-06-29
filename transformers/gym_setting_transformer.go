package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformGymSetting(setting models.GymSetting) dto.GymSettingResponse {
	return dto.GymSettingResponse{
		ID:            setting.Id,
		Gym: dto.GymResponseMin{
			ID:   setting.Gym.Id,
			Name: setting.Gym.Name,
		},
		CancelTimeLimitMinutes: setting.CancelTimeLimitMinutes,
	}
}