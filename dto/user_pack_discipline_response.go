package dto

type UserPackDisciplineResponse struct {
	ID             uint               `json:"id"`
	UserPack       UserPackResponseMin`json:"user_pack"`
	Discipline     DisciplineResponse `json:"discipline"`
}