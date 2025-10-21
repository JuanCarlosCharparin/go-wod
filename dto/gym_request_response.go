package dto

type GymRequestResponse struct {
	ID        uint            `json:"id"`
	User      UserResponseMin `json:"user"`
	Gym       GymResponseMin  `json:"gym"`
	CreatedAt string          `json:"created_at"`
}