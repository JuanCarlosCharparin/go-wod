package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformWaitlist(waitlist models.Waitlist) dto.WaitlistResponse {
	return dto.WaitlistResponse{
		ID: waitlist.ID,
		User: dto.UserResponseMin{
			ID:       waitlist.User.Id,
			Name:     waitlist.User.Name,
			Lastname: waitlist.User.Lastname,
			DNI: 	  waitlist.User.DNI,
		},
		Class: dto.ClassResponse{
			ID:   waitlist.Class.Id,
			Date: waitlist.Class.Date,
			Time: waitlist.Class.Time,
			Capacity: waitlist.Class.Capacity,
			Gym: dto.GymResponseMin{
				ID:   waitlist.Class.Gym.Id,
				Name: waitlist.Class.Gym.Name,
			},
			Discipline: dto.DisciplineResponse{
				ID:   waitlist.Class.Discipline.Id,
				Name: waitlist.Class.Discipline.Name,
			},
		},
	}
}