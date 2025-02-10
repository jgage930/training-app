package api

import (
	"log"
	"net/http"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s - %s", r.Method, r.URL, r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}
