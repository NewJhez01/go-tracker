package routing

import (
	"net/http"

	"github.com/NewJhez01/go-tracker/internal/handlers"
)

func RegisterEvents(m *http.ServeMux, h *handlers.AppRoutes) {
	m.Handle("POST /events/{user_id}", h.CreateEvents)
}
