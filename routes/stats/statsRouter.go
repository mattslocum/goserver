package routes

import (
	"fmt"
	"net/http"
)

type StatsRouter struct {}

func (h *StatsRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Stats")
}
