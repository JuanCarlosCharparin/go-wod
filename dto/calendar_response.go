package dto

type CalendarResponse struct {
	ID    uint            `json:"id"`
	User  UserResponseMin `json:"user"`
	Class ClassResponse   `json:"class"`
	PackUsage *PackUsageResponse `json:"pack_usage,omitempty"`
}

type PackUsageResponse struct {
	PackID        uint `json:"pack_id"`
	ClassQuantity int  `json:"class_quantity"`
	Used          int  `json:"used"`
	Remaining     int  `json:"remaining"`
}