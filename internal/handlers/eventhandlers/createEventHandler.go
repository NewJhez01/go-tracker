package eventhandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NewJhez01/go-tracker/internal/app/command"
)

type CreateNewEventHandler struct{}

type createNewEventHandlerBody struct {
	Headline    string  `json:"headline"`
	Description *string `json:"description"`
}

func (c CreateNewEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")

	if userID == "" {
		http.Error(w, "failed to fetch user from path", http.StatusBadRequest)
		return
	}

	intID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "failed to parse id from request", http.StatusBadRequest)
		return
	}

	b := createNewEventHandlerBody{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "request body incorrect", http.StatusBadRequest)
		return
	}

	d := command.CreateNewEventDto{
		UserID:      intID,
		Headline:    b.Headline,
		Description: b.Description,
	}

	if err := command.PersistNewEvent(&d); err != nil {
		http.Error(w, "failed to persist the event", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
