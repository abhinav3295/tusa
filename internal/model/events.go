package model

import (
	"time"

	"github.com/google/uuid"
)

type TusaEvent struct {
	Id          uuid.UUID    `json:"id"`
	EventDate   time.Time    `json:"date"`
	Venue       string       `json:"venue"`
	Category    TusaCategory `json:"category"`
	Description string       `json:"description"`
}
