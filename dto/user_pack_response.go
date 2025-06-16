package dto

type UserPackResponse struct {
	ID             uint              `json:"id"`
	StartDate      string            `json:"start_date"`       
	ExpirationDate string            `json:"expiration_date"`  
	Status         int               `json:"status"`
	Gym            GymResponseMin    `json:"gym"`
	User           UserResponseMin   `json:"user"`
	Pack           PackResponseMin   `json:"pack"`
	Discipline     DisciplineResponse`json:"discipline"`
}


type UserPackResponseMin struct {
	ID             uint              `json:"id"`
	StartDate      string            `json:"start_date"`      
	ExpirationDate string            `json:"expiration_date"`  
	Status         int               `json:"status"`
}