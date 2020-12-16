package middleware

import (
	"log"
	"net/http"
	"time"
)

var Timers = make(map[string]*HttpTimer)


func HttpTimingMiddleware(key string, next http.Handler) http.Handler {
	Timers[key] = &HttpTimer{Count:0, Duration:0}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tStart := time.Now()
		next.ServeHTTP(w, r)
		tEnd := time.Now()
		Timers[key].AddTime(int64(tEnd.Sub(tStart)))
		log.Printf("API Call took: %d", tEnd.Sub(tStart))
	})
}

type HttpTimer struct {
	Count int
	Duration int64
}

func (ht *HttpTimer) AddTime(dur int64) {
	ht.Count++
	ht.Duration += dur
}
