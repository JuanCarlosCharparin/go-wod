package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformUserPack(up models.UserPack) dto.UserPackResponse {
	// Transformar disciplinas para traer mas de una
	disciplines := make([]dto.DisciplineResponse, 0)
	for _, upd := range up.Disciplines {
		disciplines = append(disciplines, dto.DisciplineResponse{
			ID:   upd.Discipline.Id,
			Name: upd.Discipline.Name,
		})
	}

	// Si user_pack.class_quantity es nil → usar la del pack
	var effectiveClassQuantity *int
	if up.ClassQuantity != nil {
		effectiveClassQuantity = up.ClassQuantity
	} else {
		effectiveClassQuantity = &up.Pack.ClassQuantity
	}

	// Preparar respuesta mínima del pack
	var packResp dto.PackResponseMin
	if up.Pack != nil && up.Pack.Id != 0 { // 👈 si realmente está cargado
		packResp = dto.PackResponseMin{
			ID:            up.Pack.Id,
			PackName:      up.Pack.PackName,
			Price:         up.Pack.Price,
			ClassQuantity: up.Pack.ClassQuantity,
		}
	}

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
			ID:       up.User.Id,
			Name:     up.User.Name,
			Lastname: up.User.Lastname,
			DNI:      up.User.DNI,
		},
		Pack:          packResp,
		Disciplines:   disciplines,
		ClassQuantity: effectiveClassQuantity, // 👈 ya no queda null
	}
}