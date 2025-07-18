package dto

type ClassResponse struct {
	ID         uint               `json:"id"`
	Date       string             `json:"date"`
	Time       string             `json:"time"`
	Capacity   int                `json:"capacity"`
	/*Enrolled   int                `json:"enrolled"` // cantidad de usuarios anotados
	Vacancy    int                `json:"vacancy"`*/  // capacidad - inscriptos
	Gym        GymResponseMin     `json:"gym"`
	Discipline DisciplineResponse `json:"discipline"`
}


type ClassResponseInfo struct {
	ID         uint               `json:"id"`
	Date       string             `json:"date"`
	Time       string             `json:"time"`
	Capacity   int                `json:"capacity"`
	Enrolled   int                `json:"enrolled"` // cantidad de usuarios anotados
	Vacancy    int                `json:"vacancy"`  // capacidad - inscriptos
	Gym        GymResponseMin     `json:"gym"`
	Discipline DisciplineResponse `json:"discipline"`
}


type ClassWithStatusResponse struct {
	ID         uint               `json:"id"`
	Date       string             `json:"date"`
	Time       string             `json:"time"`
	Gym        GymResponseMin     `json:"gym"`
	Discipline DisciplineResponse `json:"discipline"`
	Status     string             `json:"status"`
}