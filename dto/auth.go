package dto

type RegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Lastname   string `json:"lastname" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	DNI        *string `json:"dni" binding:"required"`
	BirthDate  *string `json:"birth_date" binding:"required"` 
	Password   string `json:"password" binding:"required,min=6"`
	GymId      uint   `json:"gym_id" binding:"required"`
}


type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}