package eventhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/NewJhez01/go-tracker/internal/app/command"
	"github.com/NewJhez01/go-tracker/internal/handlers/helpers"
)

type CreateNewEventHandler struct{}

type createNewEventHandlerBody struct {
	Headline    string  `json:"headline"`
	Description *string `json:"description"`
}

func (c CreateNewEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("user_id")

	intId, err := helpers.ParseId(userId)
	if err != nil {
		http.Error(w, "invalid user id given", http.StatusBadRequest)
		return
	}

	b := createNewEventHandlerBody{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "request body incorrect", http.StatusBadRequest)
		return
	}

	d := command.CreateNewEventDto{
		UserId:      intId,
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
