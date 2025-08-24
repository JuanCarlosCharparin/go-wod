package transformers

import (
	"time"
	"wod-go/dto"
	"wod-go/models"
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


func TransformClassInfo(class models.Class, enrolled int) dto.ClassResponseInfo {
	vacancy := class.Capacity - enrolled

	dateParsed, _ := time.Parse(time.RFC3339, class.Date)
    timeParsed, _ := time.Parse("15:04:05", class.Time)

    classDateTime := time.Date(
        dateParsed.Year(), dateParsed.Month(), dateParsed.Day(),
        timeParsed.Hour(), timeParsed.Minute(), timeParsed.Second(),
        0, dateParsed.Location(),
    )

    dias := map[time.Weekday]string{
        time.Monday:    "Lunes",
        time.Tuesday:   "Martes",
        time.Wednesday: "Miércoles",
        time.Thursday:  "Jueves",
        time.Friday:    "Viernes",
        time.Saturday:  "Sábado",
        time.Sunday:    "Domingo",
    }

	return dto.ClassResponseInfo{
		ID:       class.Id,
		Date:     class.Date,
		Time:     class.Time,
		DayOfWeek:  dias[classDateTime.Weekday()],
		Capacity: class.Capacity,
		Enrolled: enrolled,
		Vacancy:  vacancy,
		Gym: dto.GymResponseMin{
			ID:   class.Gym.Id,
			Name: class.Gym.Name,
		},
		Discipline: dto.DisciplineResponse{
			ID:   class.Discipline.Id,
			Name: class.Discipline.Name,
		},
	}
}



func TransformClassWithStatus(class models.Class, status string, reserved time.Time) dto.ClassWithStatusResponse {
	return dto.ClassWithStatusResponse{
		ID:     class.Id,
		Date:   class.Date,
		Time:   class.Time,
		Gym: dto.GymResponseMin{
			ID:   class.Gym.Id,
			Name: class.Gym.Name,
		},
		Discipline: dto.DisciplineResponse{
			ID:   class.Discipline.Id,
			Name: class.Discipline.Name,
		},
		Status:  status,
		Reserved: reserved,
	}
}