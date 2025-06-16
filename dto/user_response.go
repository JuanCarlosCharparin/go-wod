package dto

type UserResponse struct {
	ID        uint        `json:"id"`
	Name      string      `json:"name"`
	Lastname  string      `json:"lastname"`
	Gender    string      `json:"gender"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	DNI       string      `json:"dni"`
	BirthDate *string     `json:"birth_date,omitempty"`
	Gym       GymResponseMin `json:"gym,omitempty"` 
}


type UserResponseMin struct {
	ID        uint        `json:"id"`
	Name      string      `json:"name"`
	Lastname  string      `json:"lastname"`
	DNI       *string      `json:"dni"`
}
