package dto

type GymSettingResponse struct {
	ID                     uint            `json:"id"`
	Gym                    GymResponseMin  `json:"gym"`
	CancelTimeLimitMinutes int		       `json:"cancel_time_limit_minutes"`
}