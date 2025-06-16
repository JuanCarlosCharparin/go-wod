package dto

type PackResponse struct {
	ID            uint           `json:"id"`
	PackName      string         `json:"pack_name"`
	Price         float64        `json:"price"`
	ClassQuantity int            `json:"class_quantity"`
	Months        int            `json:"months"`
	Gym           GymResponseMin `json:"gym"`
}


type PackResponseMin struct {
	ID            uint           `json:"id"`
	PackName      string         `json:"pack_name"`
	Price         float64        `json:"price"`
	ClassQuantity int            `json:"class_quantity"`
}