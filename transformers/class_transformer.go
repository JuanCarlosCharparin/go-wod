package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformClass(class models.Class) dto.ClassResponse {

	return dto.ClassResponse{
		ID: class.Id,
		Date: class.Date,
		Time: class.Time,
		Capacity: class.Capacity,
		Gym: dto.GymResponseMin{
			ID:   class.GymId,
			Name: class.Gym.Name,
		},
		Discipline: dto.DisciplineResponse{
			ID:   class.DisciplineId,
			Name: class.Discipline.Name,
		},
	}
}