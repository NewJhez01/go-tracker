package domain

type Event struct {
	UserID      int64
	Headline    string
	Description string
}

type EventWithoutDescription struct {
	UserID   int64
	Headline string
}

func CreateEvent(userID int64, headline, description string) Event {
	return Event{
		userID,
		headline,
		description,
	}
}

func CreateEventWithoutDescription(userID int64, headline string) EventWithoutDescription {
	return EventWithoutDescription{
		userID,
		headline,
	}
}
