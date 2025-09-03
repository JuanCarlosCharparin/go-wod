package dto

import (
	"time"
)

type CalendarResponse struct {
	ID    uint            `json:"id"`
	User  UserResponseMin `json:"user"`
	Class ClassResponse   `json:"class"`
	Status string         `json:"status"`
	PackUsage *PackUsageResponse `json:"pack_usage,omitempty"`
	Reserved   time.Time         `json:"reserved"`
	Canceled  *time.Time        `json:"canceled,omitempty"`
}

type PackUsageResponse struct {
	PackID        uint `json:"pack_id"`
	ClassQuantity int  `json:"class_quantity"`
	Used          int  `json:"used"`
	Remaining     int  `json:"remaining"`
}