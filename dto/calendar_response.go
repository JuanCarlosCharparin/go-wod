package dto

type CalendarResponse struct {
	ID    uint            `json:"id"`
	User  UserResponseMin `json:"user"`
	Class ClassResponse   `json:"class"`
}