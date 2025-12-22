package main

import (
	"net/http"

	"github.com/NewJhez01/go-tracker/cmd/routing"
	"github.com/NewJhez01/go-tracker/internal/handlers"
	"github.com/NewJhez01/go-tracker/internal/handlers/eventhandlers"
)

func main() {
	mux := http.NewServeMux()
	appRouter := handlers.AppRoutes{
		eventhandlers.CreateNewEventHandler{},
	}
	routing.RegisterEvents(mux, &appRouter)
}
