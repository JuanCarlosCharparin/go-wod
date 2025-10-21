package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformGymRequest(gym_requests models.GymRequest) dto.GymRequestResponse {
	return dto.GymRequestResponse{
		ID:   gym_requests.Id,
		User: dto.UserResponseMin{
			ID:       gym_requests.User.Id,
			Name:     gym_requests.User.Name,
			Lastname: gym_requests.User.Lastname,
			DNI: 	  gym_requests.User.DNI,
		},
		Gym:  dto.GymResponseMin{
			ID:   gym_requests.Gym.Id,
			Name: gym_requests.Gym.Name,
		},
	}
}