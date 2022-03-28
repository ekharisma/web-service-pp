package model

import "time"

type Temperature struct {
	Timestamp   time.Time
	Temperature [2]float32
}
