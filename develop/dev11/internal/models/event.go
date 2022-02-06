package models

import (
	"time"
	"log"
)
type Event struct {
	EventID		int			`json:"event_id"`
	Date		Time		`json:"date"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
}

type Time struct {
    time.Time
}

const TimeFormat = "02-01-2006"

func (t *Time) UnmarshalJSON(b []byte) error {
    
	newTime, err := time.Parse("\""+TimeFormat+"\"", string(b))
	log.Printf("%q\n",string(b))

	if err != nil {
		*t = Time{time.Now()}
		return err
	}


	*t = Time{newTime}
    return nil
}