package events

import (
	"time"
)

type TusaEvent struct {
	EventDate   time.Time
	Venue       string
	Category    TusaCategory
	Description string
}

type TusaCategory int64

const (
	BoardGame TusaCategory = iota
	Karaoke
	Clubbing
	BarEvening
	BBQ
	Sports
	Picnic
	Other
)
