package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformDiscipline(discipline models.Discipline) dto.DisciplineResponse {

	return dto.DisciplineResponse{
		ID: discipline.Id,
		Name: discipline.Name,
	}
}