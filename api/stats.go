package api

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type StatsHandler struct {
	DB *sqlx.DB
}

type ActivityStats struct {
	TotalDistance float64 `db:"total_distance" json:"total_distance"`
	AverageSpeed  float64 `db:"average_speed" json:"average_speed"`
	MaxSpeed      float64 `db:"max_speed" json:"max_speed"`
	AvgHeartRate  uint8   `db:"avg_heart_rate" json:"avg_heart_rate"`
}

func StatsRouter(h *StatsHandler, mux *http.ServeMux) {
	mux.HandleFunc("GET /stats/{id}", h.getActivityStatsById)
}

func (h *StatsHandler) getActivityStatsById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("Calculating stats for %v...", id)

	var stats ActivityStats
	h.DB.Get(
		&stats,
		`
      SELECT 
	      Max(distance) total_distance,
	      AVG(speed) average_speed,
	      Max(speed) max_speed,
	      AVG(heart_rate) average_heart_rate
      FROM activity_messages
      WHERE activity_id = $1
    `,
		id,
	)

	Respond(stats, w)
}
