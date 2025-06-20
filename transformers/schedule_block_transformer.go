package transformers

import (
	"wod-go/models"
	"wod-go/dto"
	//"time"
)

func TransformScheduleBlock(block models.ScheduleBlock) dto.ScheduleBlockResponse {
	return dto.ScheduleBlockResponse{
		ID:        block.ID,
		StartTime: block.StartTime,
		EndTime:   block.EndTime,
		Capacity:  block.Capacity,
		Discipline: dto.DisciplineResponse{
			ID:   block.Discipline.Id,
			Name: block.Discipline.Name,
		},
	}
}
