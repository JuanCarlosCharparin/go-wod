package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformUser(user models.User) dto.UserResponse {

	return dto.UserResponse{
		ID: user.Id,
		Name: user.Name,
		Lastname: user.Lastname,
		Gender: user.Gender,
		Phone: user.Phone,
		Email: user.Email,
		BirthDate: user.BirthDate,
		DNI: *user.DNI,
		Gym: dto.GymResponseMin{
			ID:   user.GymId,
			Name: user.Gym.Name,
		},
		Role: dto.RoleResponse{
			Id:   user.RoleId,
			Name: user.Role.Name,
		},
	}
}