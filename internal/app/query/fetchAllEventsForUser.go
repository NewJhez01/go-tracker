package query

import "github.com/NewJhez01/go-tracker/internal/repository"

func Handle(userId int64) string {
	repo := repository.EventsRepoStruct{}

	return repo.FetchAllEventsForUser(userId)
}
