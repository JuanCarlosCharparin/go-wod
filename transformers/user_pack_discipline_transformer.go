package transformers

import (
    "wod-go/models"
    "wod-go/dto"
)

func TransformUserPackDiscipline(upd models.UserPackDiscipline) dto.UserPackDisciplineResponse {
	return dto.UserPackDisciplineResponse{
		ID: upd.Id,
		UserPack: dto.UserPackResponseMin{
			ID: upd.UserPack.Id,
			StartDate: upd.UserPack.StartDate.Format("2006-01-02"),
			ExpirationDate: upd.UserPack.ExpirationDate.Format("2006-01-02"),
			Status: upd.UserPack.Status,
		},
		Discipline: dto.DisciplineResponse{
			ID:   upd.Discipline.Id,
			Name: upd.Discipline.Name,
		},
	}
}
