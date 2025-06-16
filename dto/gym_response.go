package dto

type GymResponse struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Location string          `json:"location"`
	Phone    string          `json:"phone"`
	Email    string          `json:"email"`
	Country  CountryResponse `json:"country"`
}

type GymResponseMin struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
}
