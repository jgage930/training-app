package api

import (
	"github.com/rs/cors"
	"log"
	"net/http"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s - %s", r.Method, r.URL, r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func CorsMiddleware(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Handler(h)
		h.ServeHTTP(w, r)
	})
}
