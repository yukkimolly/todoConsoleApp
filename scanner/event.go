package scanner

import "time"

type Event struct {
	Name      string
	UserInput string

	DateTime time.Time
}

func newEvent(name string, userInput string) *Event {
	return &Event{
		Name:      name,
		UserInput: userInput,

		DateTime: time.Now(),
	}
}
