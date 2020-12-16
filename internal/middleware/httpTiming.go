package middleware

import (
	"net/http"
	"sync"
	"time"
)

var Timers = make(map[string]*EventTimer)
var mu sync.Mutex // lock for timer creation

func HttpTimingMiddleware(key string, next http.Handler) http.Handler {
	EnsureTimer(key)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tStart := time.Now()
		next.ServeHTTP(w, r)
		tEnd := time.Now()
		Timers[key].AddEvent(int64(tEnd.Sub(tStart)))
	})
}

func EnsureTimer(key string) {
	mu.Lock()
	if Timers[key] == nil {
		Timers[key] = &EventTimer{Count: 0, Duration:0}
	}
	mu.Unlock()
}

type EventTimer struct {
	Count int
	Duration int64
}

func (ht *EventTimer) AddEvent(dur int64) {
	ht.Count++
	ht.Duration += dur
}

// FUTURE: Consider having startTimer and endTimer and maybe have startTimer make a
//         chan to do timing instead of making the integrator run their own timing.
