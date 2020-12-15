package routes

import (
	"fmt"
	"net/http"
)

type StatsHandler struct {}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Stats")
}
