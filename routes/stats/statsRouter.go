package routes

import (
	"encoding/json"
	"github.com/mattslocum/goserver/internal/middleware"
	"math"
	"net/http"
)

type statsData struct {
	Total int `json:"total"`
	Average float64 `json:"average"`
}

type StatsRouter struct {}

func (h *StatsRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	timer := middleware.Timers["GetHash"]
	average := 0.0
	if timer.Count > 0 {
		average = math.Round(float64(timer.Duration) / float64(timer.Count))
	}

	data := statsData{
		Total:   timer.Count,
		Average: average,
	}
	json.NewEncoder(w).Encode(data)
}
