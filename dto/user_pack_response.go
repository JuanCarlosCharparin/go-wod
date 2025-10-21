package dto

import "time"

type UserPackResponse struct {
	ID             uint               `json:"id"`
	StartDate      time.Time          `json:"start_date"`
	ExpirationDate time.Time          `json:"expiration_date"`
	Status         int                `json:"status"`
	Gym            GymResponseMin     `json:"gym"`
	User           UserResponseMin    `json:"user"`
	Pack           PackResponseMin    `json:"pack"`
	ClassQuantity  *int               `json:"class_quantity"`
	Disciplines    []DisciplineResponse `json:"disciplines"`
}

type UserPackUsageItem struct {
	UserPackID     uint            `json:"user_pack_id"`
	Status         int             `json:"status"` // 1 activo, 0 vencido
	StartDate      string          `json:"start_date"`
	ExpirationDate string          `json:"expiration_date"`
	Used           int             `json:"used"`
	Remaining      int             `json:"remaining"`
	ClassQuantity  int             `json:"class_quantity"`
	Pack           *PackResponseMin `json:"pack,omitempty"`
	Disciplines    []DisciplineResponse `json:"disciplines"`
	Gym            GymResponseMin  `json:"gym"`
}

type UserPackResponseMin struct {
	ID             uint   `json:"id"`
	StartDate      string `json:"start_date"`
	ExpirationDate string `json:"expiration_date"`
	Status         int    `json:"status"`
}