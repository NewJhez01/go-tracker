package eventhandlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("failed to fetch user id from path")
	}

	intID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Printf("failed to parse int from user id: %s", userID)
	}

	b := createNewEventHandlerBody{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Println("work in progress")
	}

	d := command.CreateNewEventDto{
		UserID:      intID,
		Headline:    b.Headline,
		Description: b.Description,
	}

	if err := command.PersistNewEvent(&d); err != nil {
		fmt.Println("work in progress")
	}

	fmt.Println("success")
}
