package model

import "github.com/google/uuid"

type RSVP struct {
	Id             uuid.UUID `json:"id"`
	EventId        uuid.UUID `json:"eventId"`
	UserId         uuid.UUID `json:"userId"`
	Attending      bool      `json:"attending"`
	NumberOfGuests int       `json:"numberOfGuests"`
}
