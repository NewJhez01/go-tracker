package repository

import (
	"errors"

	domain "github.com/NewJhez01/go-tracker/internal/domain/event"
)

type eventsRepo interface {
	persistNewEvent(*domain.Event) error
	persistNewEventWithoutDescription(*domain.Event) error
}

type eventsRepoStruct struct{}

func (e eventsRepoStruct) persistNewEvent(d *domain.Event) error {
	return errors.New("work in progress")
}

func (e eventsRepoStruct) persistNewEventWithoutDescription(d *domain.Event) error {
	return errors.New("work in progress")
}
