package dto

type WodResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
	Duration    int             `json:"duration"`
	Level       string          `json:"level"`
	Gym         GymResponseMin		`json:"gym"`
}