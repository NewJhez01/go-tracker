package repository

import (
	"fmt"

	domain "github.com/NewJhez01/go-tracker/internal/domain/event"
)

type EventsRepoStruct struct{}

func (e EventsRepoStruct) PersistNewEvent(d *domain.Event) {
	s := fmt.Sprintf("endpoint hit successfully with data headline: %s and desc: %s", d.Headline, d.Description)
	fmt.Println(s)
}

func (e EventsRepoStruct) PersistNewEventWithoutDescription(d *domain.EventWithoutDescription) {
	s := fmt.Sprintf("endpoint hit successfully with data with headline %s ", d.Headline)
	fmt.Println(s)
}

func (e EventsRepoStruct) FetchAllEventsForUser(userId int64) string {
	return "hello world"
}
