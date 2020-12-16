package routes

import (
	"github.com/mattslocum/goserver/internal/middleware"
	hash "github.com/mattslocum/goserver/routes/hash"
	shutdown "github.com/mattslocum/goserver/routes/shutdown"
	stats "github.com/mattslocum/goserver/routes/stats"
	"log"
	"net/http"
)

var logger = middleware.HttpLoggingMiddleware

func setupRoutes() {
	// TODO: Error handler and logger
	// better pattern matching?
	http.Handle("/hash", logger(middleware.HttpTimingMiddleware("GetHash", hash.GetHashRouter())))
	http.Handle("/hash/", logger(hash.GetHashRouter()))
	http.Handle("/shutdown", logger(new(shutdown.ShutdownRouter)))
	http.Handle("/stats", logger(new(stats.StatsRouter)))
}

func Setup() {
	setupRoutes()

	log.Print("Starting Server on port 8080")
	// Do we need to do http.Server?
	log.Fatal(http.ListenAndServe(":8080", nil))
}

