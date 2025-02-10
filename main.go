package main

import (
	"net/http"
	"text/template"
	"training-app/api"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)

	mux.HandleFunc("/home", Index)

	server := api.LoggingMiddleware(mux)

	http.ListenAndServe(":8080", server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, "")
}
