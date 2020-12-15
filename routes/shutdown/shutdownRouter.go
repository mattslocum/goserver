package routes

import (
	"fmt"
	"net/http"
)

type ShutdownRouter struct {}

func (h *ShutdownRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Shutdown!")
}
