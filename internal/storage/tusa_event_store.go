package storage

import (
	"time"
	"tusa/internal/model"

	"github.com/google/uuid"
)

type TusaEventStore interface {
	Add(event model.TusaEvent) error
	Delete(event model.TusaEvent) error
	Find(id uuid.UUID) *model.TusaEvent
	FindAllLatest(since time.Time) []model.TusaEvent
}

type inMemoryTusaEventStore struct {
	store map[uuid.UUID]model.TusaEvent
}

func NewTusaEventStore() inMemoryTusaEventStore {
	return inMemoryTusaEventStore{
		store: make(map[uuid.UUID]model.TusaEvent),
	}
}

func (s *inMemoryTusaEventStore) Add(event model.TusaEvent) error {
	s.store[event.Id] = event
	return nil
}
func (s *inMemoryTusaEventStore) Delete(event model.TusaEvent) error {
	delete(s.store, event.Id)
	return nil
}
func (s *inMemoryTusaEventStore) Find(id uuid.UUID) *model.TusaEvent {
	event, found := s.store[id]
	if found {
		return &event
	}
	return nil
}
func (s *inMemoryTusaEventStore) FindAllLatest(since time.Time) []model.TusaEvent {
	values := make([]model.TusaEvent, 0, len(s.store))
	for _, val := range s.store {
		values = append(values, val)
	}
	return values
}
