package models

type GymSetting struct {
	Id              uint        `json:"id" gorm:"primaryKey"`
	GymId           uint        `json:"gym_id"`           
	Gym             Gym         `json:"gym" gorm:"foreignKey:GymId"`
	CancelTimeLimitMinutes int  `json:"cancel_time_limit_minutes"`
}
