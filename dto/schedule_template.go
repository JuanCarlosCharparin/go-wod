package dto


type ScheduleTemplateResponse struct {
	ID             uint                    `json:"id"`
	Day      	   string                  `json:"day"`       
	Gym            GymResponseMin          `json:"gym"`
	Blocks         []ScheduleBlockResponse `json:"blocks"`
}