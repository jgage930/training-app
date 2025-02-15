package api

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type StatsHandler struct {
	DB *sqlx.DB
}

func StatsRouter(h *StatsHandler, mux *http.ServeMux) {
	mux.HandleFunc("GET /stats/{name}/{id}", h.getActivityStatsById)
}

func (h *StatsHandler) getActivityStatsById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	name := r.PathValue("name")

	log.Printf("Calculating %v stats for %v...", name, id)

	var stats ActivityStats
	switch name {
	case "activity":
		stats = activityStatsById(h.DB, id)
	default:
		http.Error(w, "Stat type not supported", http.StatusBadRequest)
	}

	Convert(&stats)
	Respond(stats, w)
}

type Stats interface{}

type ActivityStats struct {
	TotalDistance float64 `db:"total_distance" json:"total_distance"`
	AverageSpeed  float64 `db:"average_speed" json:"average_speed"`
	MaxSpeed      float64 `db:"max_speed" json:"max_speed"`
	AvgHeartRate  uint8   `db:"average_heart_rate" json:"average_heart_rate"`
}

func Convert(stats *ActivityStats) {
	stats.TotalDistance = ConvertValue(stats.TotalDistance, Meter, Mile)
	stats.AverageSpeed = ConvertValue(stats.AverageSpeed, MetersPerSecond, MilesPerHour)
	stats.MaxSpeed = ConvertValue(stats.MaxSpeed, MetersPerSecond, MilesPerHour)
}

func activityStatsById(db *sqlx.DB, id string) ActivityStats {
	var stats ActivityStats
	db.Get(
		&stats,
		`
	     SELECT
	      Max(distance) total_distance,
	      Avg(speed) average_speed,
	      Max(speed) max_speed,
	      Avg(heart_rate) average_heart_rate
	     FROM activity_messages
	     WHERE activity_id = $1
	   `,
		id,
	)

	return stats
}
