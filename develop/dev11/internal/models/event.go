package models

import (
	"time"
	"log"
)

//Event represents business entity "Event"
type Event struct {
	EventID		int			`json:"event_id"`
	Date		Time		`json:"date"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
}

//Time is a struct for marshal/unmarshal time
type Time struct {
	time.Time
}

//TimeFormat is a layout for representing time
const TimeFormat = "02-01-2006"

//UnmarshalJSON is made for transforming json to Time struct
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