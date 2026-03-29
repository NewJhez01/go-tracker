package main

import (
	"fmt"
	"net/http"

	"github.com/NewJhez01/go-tracker/cmd/routing"
	"github.com/NewJhez01/go-tracker/internal/handlers"
	"github.com/NewJhez01/go-tracker/internal/handlers/eventhandlers"
)

func main() {
	mux := http.NewServeMux()
	appRouter := handlers.AppRoutes{
		CreateEvents: eventhandlers.CreateNewEventHandler{},
		ReadEvents:   eventhandlers.ReadAllEventHandler{},
	}
	routing.RegisterEventHandlers(mux, &appRouter)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Print("fatal")
	}
}
