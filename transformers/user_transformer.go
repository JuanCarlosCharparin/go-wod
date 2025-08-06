package transformers

import (
	"wod-go/dto"
	"wod-go/models"
)

func TransformUser(user models.User) dto.UserResponse {

	return dto.UserResponse{
		ID:        user.Id,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Gender:    user.Gender,
		Phone:     user.Phone,
		Email:     user.Email,
		BirthDate: user.BirthDate,
		DNI:       *user.DNI,
		Gym: dto.GymResponseMin{
			ID:   user.GymId,
			Name: user.Gym.Name,
		},
		Role: dto.RoleResponse{
			Id:   user.RoleId,
			Name: user.Role.Name,
		},
		Status: user.Status,
	}
}

func TransformUserNoGym(user models.User) dto.UserResponseNoGym {

	return dto.UserResponseNoGym{
		ID:        user.Id,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Gender:    user.Gender,
		Phone:     user.Phone,
		Email:     user.Email,
		BirthDate: user.BirthDate,
		DNI:       *user.DNI,
		Role: dto.RoleResponse{
			Id:   user.RoleId,
			Name: user.Role.Name,
		},
		Status: user.Status,
		UserPack: func() []dto.UserPackResponseMin {
			packs := make([]dto.UserPackResponseMin, 0, len(user.UserPacks))
			for _, up := range user.UserPacks {
				packs = append(packs, dto.UserPackResponseMin{
					ID:             up.Id,
					StartDate:      up.StartDate.Format("2006-01-02"),
					ExpirationDate: up.ExpirationDate.Format("2006-01-02"),
					Status:         up.Status,
				})
			}
			return packs
		}(),
	}
}
