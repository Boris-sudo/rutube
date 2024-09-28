package mw

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs entry and exit of API calls
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log entry
		start := time.Now()
		log.Printf("Started %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log exit
		log.Printf("Completed %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}
