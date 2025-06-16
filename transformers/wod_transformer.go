package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformWod(wod models.Wod) dto.WodResponse {

	return dto.WodResponse{
		ID: wod.Id,
		Name: wod.Name,
		Description: wod.Description,
		Type: wod.Type,
		Duration: wod.Duration,
		Level: wod.Level,
		Gym: dto.GymResponseMin{
			ID:   wod.GymId,
			Name: wod.Gym.Name,
		},
	}
}