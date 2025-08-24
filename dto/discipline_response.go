package dto 

type DisciplineResponse struct {
	ID   uint            `json:"id"`
	Name string          `json:"name"`
	Gym  GymResponseMin  `json:"gym"`
}