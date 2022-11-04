package tracker

import (
	"fmt"
	"mail_sender/internal/app"
	"time"
)

type Tracker struct {
	Events []Event
}

type Event struct {
	Event     app.TrackMailParam
	timestamp time.Time
}

func NewTracker() *Tracker {

	events := *new([]Event)
	return &Tracker{
		Events: events,
	}
}

func (t *Tracker) Track(param app.TrackMailParam) {
	event := Event{
		Event:     param,
		timestamp: time.Now(),
	}

	t.Events = append(t.Events, event)

	fmt.Println("----- Events ------- :    ")
	for i, event := range t.Events {
		fmt.Printf("Id: %d,	Timestamp: %s,	Event: %+v\n", i, event.timestamp, event.Event)
	}
	fmt.Println("--------------------")
}
