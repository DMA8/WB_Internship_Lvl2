package repository

import (
	"../models"
)

type LocalSotrage struct {
	Events	[]models.Event
}

func NewLocalStorage() *LocalSotrage{
	return &LocalSotrage{
		Events: make([]models.Event, 0),
	}
}

func (s *LocalSotrage)CreateEvent(event models.Event) error {
	s.Events = append(s.Events, event)
	return nil
}

func (s *LocalSotrage)GetAllEvents()([]models.Event, error) {
	out := make([]models.Event, len(s.Events))
	copy(out, s.Events)
	return out, nil
}