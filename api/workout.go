package api

import (
	"net/http"
	"time"
)

type WorkoutHandler struct{}

func WorkoutRouter(h *WorkoutHandler, mux *http.ServeMux) {
	mux.HandleFunc("POST /workout", h.postWorkout)
}

type WorkoutCreate struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

func (h *WorkoutHandler) postWorkout(w http.ResponseWriter, r *http.Request) {
	var payload WorkoutCreate
	ReadBody(&payload, w, r)

	// Create workout in db.
}
