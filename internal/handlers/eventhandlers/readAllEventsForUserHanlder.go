package eventhandlers

import "net/http"

type ReadAllEventHandler struct{}

func (ra ReadAllEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
