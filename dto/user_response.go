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
	Gym       GymResponseMin `json:"gym"` 
	Role	  RoleResponse `json:"role"`
	Status	  bool        `json:"status"` 
}


type UserResponseNoGym struct {
	ID        uint        `json:"id"`
	Name      string      `json:"name"`
	Lastname  string      `json:"lastname"`
	Gender    string      `json:"gender"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	DNI       string      `json:"dni"`
	BirthDate *string     `json:"birth_date,omitempty"`
	Role	  RoleResponse `json:"role"` 
	Status	  bool        `json:"status"`
	UserPack  []UserPackResponseMin `json:"user_packs,omitempty"` // Lista de UserPackResponseMin
}


type PaginatedUsersResponse struct {
	Data       []UserResponseNoGym `json:"data"`
	Total      int64                   `json:"total"`
	Page       int                     `json:"page"`
	Limit      int                     `json:"limit"`
	TotalPages int                     `json:"total_pages"`
}


type UserResponseMin struct {
	ID        uint        `json:"id"`
	Name      string      `json:"name"`
	Lastname  string      `json:"lastname"`
	DNI       *string      `json:"dni"`
}
