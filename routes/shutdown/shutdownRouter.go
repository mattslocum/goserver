package routes

import (
	"fmt"
	"github.com/mattslocum/goserver/internal/shutdown"
	"net/http"
)

type ShutdownRouter struct {}

func (h *ShutdownRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Shutting Down")
	shutdown.Shutdown()
}
