package command

import (
	domain "github.com/NewJhez01/go-tracker/internal/domain/event"
	"github.com/NewJhez01/go-tracker/internal/repository"
)

type CreateNewEventDto struct {
	UserID      int64
	Headline    string
	Description *string
}

func PersistNewEvent(c *CreateNewEventDto) error {
	repo := repository.EventsRepoStruct{}

	if c.Description == nil {
		model := domain.CreateEventWithoutDescription(c.UserID, c.Headline)
		repo.PersistNewEventWithoutDescription(&model)
		return nil
	}

	model := domain.CreateEvent(c.UserID, c.Headline, *c.Description)
	repo.PersistNewEvent(&model)

	return nil
}
