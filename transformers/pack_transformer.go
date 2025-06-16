// transformers/pack.go
package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformPack(pack models.Pack) dto.PackResponse {
	return dto.PackResponse{
		ID:            pack.Id,
		PackName:      pack.PackName,
		Price:         pack.Price,
		ClassQuantity: pack.ClassQuantity,
		Months:        pack.Months,
		Gym: dto.GymResponseMin{
			ID:   pack.Gym.Id,
			Name: pack.Gym.Name,
		},
	}
}
