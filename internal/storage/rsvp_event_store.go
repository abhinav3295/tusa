package storage

import (
	"tusa/internal/model"

	"github.com/google/uuid"
)

type RsvpStore interface {
	AddOrUpdate(rsvp model.RSVP) error
}

func NewRsvpStore() RsvpStore {
	return &inMemoryRsvpStore{
		store: make(map[uuid.UUID]model.RSVP),
	}
}

type inMemoryRsvpStore struct {
	store map[uuid.UUID]model.RSVP
}

func (i *inMemoryRsvpStore) AddOrUpdate(rsvp model.RSVP) error {
	i.store[rsvp.Id] = rsvp
	return nil
}
