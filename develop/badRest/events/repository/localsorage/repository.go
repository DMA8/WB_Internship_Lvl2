package repository

import (
	"../../../models/models"
)
// Опишу хранилище
// map[time.Time]Event
// sorted []string{Event, Event}
type EventRepository interface{
	GetAll()([]models.Event, error)
	Add(models.Event)error
}

type RepositoryEvent struct{
	AllEvents	[]models.Event
}

func NewRepositoryEvents()*RepositoryEvent{
	return &RepositoryEvent{
		make([]models.Event, 0),
	}
}

func (r RepositoryEvent)GetAll()[]models.Event{
	return r.AllEvents
}

func (r *RepositoryEvent)Add(obj models.Event){
	r.AllEvents = append(r.AllEvents, obj)
}
