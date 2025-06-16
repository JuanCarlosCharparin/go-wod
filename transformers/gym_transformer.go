package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformGym(gym models.Gym) dto.GymResponse {

	return dto.GymResponse{
		ID: gym.Id,
		Name: gym.Name,
		Location: gym.Location,
		Phone: gym.Phone,
		Email: gym.Email,
		Country: dto.CountryResponse{
			ID:   gym.Country.Id,
			Name: gym.Country.Name,
		},
	}
}