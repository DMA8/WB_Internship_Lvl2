package usecases

import (
	"../repository"
	"../models"

)

type EventUseCase struct {
	EventRepo repository.LocalSotrage
}

func NewEventUseCase(rep repository.LocalSotrage) *EventUseCase{
	return &EventUseCase{
		EventRepo: rep,
	}
}

func (e *EventUseCase)CreateEvent(event models.Event) error{
	return e.EventRepo.CreateEvent(event)
}

func (e *EventUseCase)GetAllEvents()([]models.Event, error){
	return e.EventRepo.GetAllEvents()
}