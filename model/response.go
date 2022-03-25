package model

import "time"

type Payload struct {
	Timestamp   time.Time  `json:"timestamp"`
	Temperature [2]float32 `json:"temperature"`
}
