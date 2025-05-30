package main

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs each HTTP request with method, path, and response time
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a custom ResponseWriter to capture status code
		ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Process the request
		next.ServeHTTP(ww, r)

		// Calculate response time
		duration := time.Since(start)

		// Log the request details
		log.Printf("[%s] %s %s - Status: %d - Duration: %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			ww.statusCode,
			duration,
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code before writing the header
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
