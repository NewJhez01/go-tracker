package eventhandlers

import (
	"fmt"
	"net/http"

	"github.com/NewJhez01/go-tracker/internal/handlers/helpers"
)

type ReadAllEventHandler struct{}

func (ra ReadAllEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("user_id")

	intId, err := helpers.ParseId(userId)
	if err != nil {
		http.Error(w, "invalid user id given", http.StatusBadRequest)
	}

	fmt.Println(intId)
}
