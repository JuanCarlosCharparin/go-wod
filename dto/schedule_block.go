package dto


type ScheduleBlockResponse struct {
	ID                uint                     `json:"id"`
	//ScheduleTemplate  ScheduleTemplateResponse `json:"template_id"`      
	StartTime         string                   `json:"start_time"` 
	EndTime           string                   `json:"end_time"`   
	Capacity          int                      `json:"capacity"`
	Discipline        DisciplineResponse       `json:"discipline"`
}