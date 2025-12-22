package repository

import (
	"fmt"

	domain "github.com/NewJhez01/go-tracker/internal/domain/event"
)

type EventsRepoStruct struct{}

func (e EventsRepoStruct) PersistNewEvent(d *domain.Event) {
	fmt.Printf("endpoint hit successfully with data %s", d.Headline)
}

func (e EventsRepoStruct) PersistNewEventWithoutDescription(d *domain.EventWithoutDescription) {
	fmt.Printf("endpoint hit successfully with data %s", d.Headline)
}
