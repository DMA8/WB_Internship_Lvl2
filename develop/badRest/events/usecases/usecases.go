package usecases

import (
	"../repository/localstorage/repository"
	"../../models/models"
	
)
// опишу как добавлять/обновлять/удалять/отдавать сущности из репозитория

type UseCaser interface{
	GetAllEvents()[]models.Event
	Add(models.Event)
}


type EventUseCase struct{
	Repo	*repository.EventRepository
}


func (e *EventUseCase)GetAllEvents()[]models.Event{
	return e.Repo.GetAllEvents()
}

func (e *EventUseCase)Add(m models.Event){
	e.Add(m)
}