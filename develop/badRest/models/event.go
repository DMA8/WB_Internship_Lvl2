package models

import "time"

// Опишу объект "Событие"

type Event struct{
	Name	string
	Date	time.Time
}