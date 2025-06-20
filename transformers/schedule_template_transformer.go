package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformScheduleTemplate(template models.ScheduleTemplate) dto.ScheduleTemplateResponse {
	var blocks []dto.ScheduleBlockResponse
	for _, block := range template.Blocks {
		blocks = append(blocks, TransformScheduleBlock(block))
	}

	return dto.ScheduleTemplateResponse{
		ID:  template.ID,
		Day: template.Day,
		Gym: dto.GymResponseMin{
			ID:   template.Gym.Id,
			Name: template.Gym.Name,
		},
		Blocks: blocks,
	}
}
