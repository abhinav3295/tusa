package events

import (
	"time"
)

type TusaEvent struct {
	EventDate   time.Time    `json:"date"`
	Venue       string       `json:"venue"`
	Category    TusaCategory `json:"category"`
	Description string       `json:"description"`
}
