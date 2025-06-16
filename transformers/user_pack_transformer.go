package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformUserPack(up models.UserPack) dto.UserPackResponse {
	return dto.UserPackResponse{
		ID:             up.Id,
		StartDate:      up.StartDate,
		ExpirationDate: up.ExpirationDate,
		Status:         up.Status,
		Gym: dto.GymResponseMin{
			ID:   up.Gym.Id,
			Name: up.Gym.Name,
		},
		User: dto.UserResponseMin{
			ID:   up.User.Id,
			Name: up.User.Name,
			Lastname: up.User.Lastname,
			DNI: up.User.DNI,
		},
		Pack: dto.PackResponseMin{
			ID:   up.Pack.Id,
			PackName: up.Pack.PackName,
			Price: up.Pack.Price,
			ClassQuantity: up.Pack.ClassQuantity,
		},
		Discipline: dto.DisciplineResponse{
			ID:   up.Discipline.Id,
			Name: up.Discipline.Name,
		},
	}
}