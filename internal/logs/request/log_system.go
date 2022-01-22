package request

import "time"

type LogSystemRequest struct {
	File      string    `json:"file"`
	Level     string    `json:"Level"`
	Err       string    `json:"err"`
	Message   string    `json:"message"`
	Line      int       `json:"line"`
	OccurTime time.Time `json:"occurTime"`
}
