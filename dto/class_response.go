package dto

type ClassResponse struct {
	ID         uint               `json:"id"`
	Date       string             `json:"date"`
	Time       string             `json:"time"`
	Capacity   int                `json:"capacity"`
	Gym        GymResponseMin     `json:"gym"`
	Discipline DisciplineResponse `json:"discipline"`
}