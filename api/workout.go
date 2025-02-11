package api

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

type Workout struct {
	Id          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Date        time.Time `db:"date" json:"date"`
	Description string    `db:"description" json:"description"`
}

type WorkoutHandler struct {
	DB *sqlx.DB
}

func WorkoutRouter(h *WorkoutHandler, mux *http.ServeMux) {
	mux.HandleFunc("POST /workout", h.postWorkout)
	mux.HandleFunc("GET /workout", h.listWorkouts)
}

type WorkoutCreate struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

func (h *WorkoutHandler) postWorkout(w http.ResponseWriter, r *http.Request) {
	var payload WorkoutCreate
	ReadBody(&payload, w, r)

	tx := h.DB.MustBegin()
	tx.MustExec(
		"INSERT INTO workouts (name, date, description) VALUES ($1, $2, $3)",
		payload.Name,
		payload.Date,
		payload.Description,
	)
	tx.Commit()
}

func (h *WorkoutHandler) listWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts := []Workout{}
	h.DB.Select(&workouts, "SELECT * from workouts;")
	Respond(workouts, w)
}
