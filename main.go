package main

import (
	"net/http"
	"text/template"
	"training-app/api"
)

func main() {
	// Database
	db := api.SetupDB()

	// Server
	mux := http.NewServeMux()

	// Static files Routes
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)

	// View Routes
	mux.HandleFunc("/home", Index)

	// Api Routes
	workoutHandler := api.WorkoutHandler{DB: db}
	api.WorkoutRouter(&workoutHandler, mux)

	activityHandler := api.ActivityHandler{DB: db}
	api.ActivityRouter(&activityHandler, mux)

	statsHandler := api.StatsHandler{DB: db}
	api.StatsRouter(&statsHandler, mux)

	// Global middleware
	server := api.LoggingMiddleware(mux)
	server = api.CorsMiddleware(server)

	http.ListenAndServe(":8080", server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, "")
}
