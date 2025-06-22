package transformers

import (
	"wod-go/dto"
	"wod-go/models"
	"wod-go/services"
)

func TransformCalendar(calendar models.Calendar, usage *services.PackUsage) dto.CalendarResponse {
	var usageDto *dto.PackUsageResponse
	if usage != nil {
		usageDto = &dto.PackUsageResponse{
			PackID:        usage.PackID,
			ClassQuantity: usage.ClassQuantity,
			Used:          usage.Used,
			Remaining:     usage.Remaining,
		}
	}

	return dto.CalendarResponse{
		ID: calendar.Id,
		User: dto.UserResponseMin{
			ID:       calendar.User.Id,
			Name:     calendar.User.Name,
			Lastname: calendar.User.Lastname,
			DNI:      calendar.User.DNI,
		},
		Class: dto.ClassResponse{
			ID:       calendar.Class.Id,
			Date:     calendar.Class.Date,
			Time:     calendar.Class.Time,
			Capacity: calendar.Class.Capacity,
			Gym: dto.GymResponseMin{
				ID:   calendar.Class.Gym.Id,
				Name: calendar.Class.Gym.Name,
			},
			Discipline: dto.DisciplineResponse{
				ID:   calendar.Class.Discipline.Id,
				Name: calendar.Class.Discipline.Name,
			},
		},
		PackUsage: usageDto,
	}
}

