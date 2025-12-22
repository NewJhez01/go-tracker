package command

import domain "github.com/NewJhez01/go-tracker/internal/domain/event"

type CreateNewEventDto struct {
	UserID      int64
	Headline    string
	Description *string
}

func PersistNewEvent(c *CreateNewEventDto) error {
	if c.Description == nil {
		domain.CreateEventWithoutDescription(c.UserID, c.Headline)
	}

	domain.CreateEvent(c.UserID, c.Headline, *c.Description)

	return nil
}
