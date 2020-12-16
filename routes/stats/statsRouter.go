package routes

import (
	"encoding/json"
	"github.com/mattslocum/goserver/internal/middleware"
	routes "github.com/mattslocum/goserver/routes/hash"
	"math"
	"net/http"
)

type statsData struct {
	Total int `json:"total"`
	Average float64 `json:"average"`
}

type StatsRouter struct {}

func (h *StatsRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// For now, only give stats on hashing
	middleware.EnsureTimer(routes.HashTimerName)
	hashTimer := middleware.Timers[routes.HashTimerName]

	average := 0.0
	if hashTimer.Count > 0 {
		average = math.Round(float64(hashTimer.Duration) / float64(hashTimer.Count))
	}

	data := statsData{
		Total:   hashTimer.Count,
		Average: average,
	}
	json.NewEncoder(w).Encode(data)
}
