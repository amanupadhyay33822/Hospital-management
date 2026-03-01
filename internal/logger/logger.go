package logger

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var start time.Time
		start = time.Now()

		var rw *responseWriter
		rw = &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		var duration time.Duration
		duration = time.Since(start)

		log.Printf(
			"%s %s %d %.3f ms",
			r.Method,
			r.URL.RequestURI(), // includes query params
			rw.status,
			float64(duration.Microseconds())/1000.0,
		)
	})
}
