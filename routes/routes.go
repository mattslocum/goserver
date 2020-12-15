package routes

import (
	"fmt"
	hash "github.com/mattslocum/goserver/routes/hash"
	shutdown "github.com/mattslocum/goserver/routes/shutdown"
	stats "github.com/mattslocum/goserver/routes/stats"
	"log"
	"net/http"
)

func baseRoute(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: homePage")
}

func setupRoutes() {
	// TODO: Error handler and logger
	//http.HandleFunc("/", baseRoute)
	// better pattern matching?
	http.Handle("/hash", hash.NewHashRouter())
	http.Handle("/hash/", hash.NewHashRouter())
	http.Handle("/shutdown", new(shutdown.ShutdownRouter))
	http.Handle("/stats", new(stats.StatsRouter))
}

func Setup() {
	setupRoutes()

	// Do we need to do http.Server?
	log.Fatal(http.ListenAndServe(":8080", nil))
}

