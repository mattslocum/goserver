package routes

import (
	"fmt"
	"net/http"
)

type ShutdownHandler struct {}

func (h *ShutdownHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Shutdown!")
}
