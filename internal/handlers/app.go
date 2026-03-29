package handlers

import (
	"net/http"
)

type AppRoutes struct {
	CreateEvents http.Handler
	ReadEvents   http.Handler
}
