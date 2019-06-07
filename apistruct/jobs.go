package apistruct

import "time"

type ReturnJobData struct {
	Id     string `json:"id,omitempty"`
	Worker string `json:"worker,omitempty"`
	Data   string `json:"data,omitempty"`
	ScheduledTime time.Time `json:"scheduledtime,omitempty"`
	CompletedTime time.Time `json:"completedtime,omitempty"`
}
