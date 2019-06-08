package handlers

import "time"

var NilTime = time.Time{}

type Status struct {
	State string `json:"state,omitempty"`
}


