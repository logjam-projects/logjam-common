package apistruct

import "time"

type CreateJob struct {
	TaskType     int      `json:"tasktype,omitempty"`
	Description  string   `json:"description,omitempty"`
	TargetWorker string   `json:"worker,omitempty"`
	TargetTags   []string `json:"tags,omitempty"`
}

type ReturnJobData struct {
	Id     string `json:"id,omitempty"`
	Worker string `json:"worker,omitempty"`
	Data   string `json:"data,omitempty"`
	ScheduledTime time.Time `json:"scheduledtime,omitempty"`
	CompletedTime time.Time `json:"completedtime,omitempty"`
	Duration 	  time.Duration `json:"duration,omitempty"`
}

type JobDetails struct {
	Id           string `json:"id,omitempty"`
	Description  string `json:"description,omitempty"`
	TaskType     int    `json:"tasktype,omitempty"`
	TargetWorker string `json:"worker,omitempty"`
	CreateTime     time.Time `json:"createtime,omitempty"`
	PulledTime     time.Time `json:"pulledtime,omitempty"`
	CompletedTime  time.Time `json:"completedtime,omitempty"`
}