package storage

import (
	"errors"
	"fmt"
	"time"

	"../models"
)

// хранить события в мапе, а отсортированные ключи мапы в срезе
// когда нам надо будет отдать месячные события, мы идем по отсорт. срезу ключей
// проверяем, в пределах ли месяца ивент. Как только мы встречаем ивент за пределами месяца,
// мы прекращаем поиск и отдаем срез ивентов.

type EventStorage struct {
	ID			int
	Events		map[int]models.Event
	SortedKeys	[]int
}

func NewEventStorage() *EventStorage{
	return &EventStorage{
		ID: 0,
		Events: make(map[int]models.Event),
	}
}

func (s *EventStorage) CreateEvent(event models.Event) (models.Event, error) {
	ID := s.ID + 1
	s.ID++
	s.Events[ID] = event
	return event, nil
}

func (s *EventStorage) UpdateEvent(ID int, updatedEvent models.Event) (models.Event, error) {
	event, ok := s.Events[ID]
	if !ok {
		return models.Event{}, errors.New("event not found")
	}
	s.Events[ID] = updatedEvent
	return event, nil
}

func (s *EventStorage) DeleteEvent(ID int) error{
	_, ok := s.Events[ID]
	if !ok {
		return errors.New("can't delete! Event doesn't exist")
	}
	delete(s.Events, ID)
	fmt.Println(s.Events)
	return nil
}


func (s *EventStorage) GetEventsForDay(date time.Time) []models.Event{
	res := []models.Event{}
	y1, m1, d1 :=date.Date()

	for _, event := range s.Events{
		y2, m2, d2 := event.Date.Date()
		if y1 == y2 && m1 == m2 && d1 == d2{
			res = append(res, event)
		}
	}
	return res
}

func (s *EventStorage) GetEventsForWeek(date time.Time)[]models.Event{
	res := []models.Event{}
	y1, w1 := date.ISOWeek()
	for _, event := range s.Events{
		y2, w2 := event.Date.ISOWeek()
		if y1 == y2 && w1 == w2{
			res = append(res, event)
		}
	}
	return res
}

func (s *EventStorage) GetEventsForMonth(date time.Time)[]models.Event{
	res := []models.Event{}
	y1, m1, _ := date.Date()
	for _, event := range s.Events{
		y2, m2, _ := event.Date.Date()
		if y1 == y2 && m1 == m2{
			res = append(res, event)
		}
	}
	return res
}