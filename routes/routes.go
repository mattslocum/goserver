package routes

import (
	"context"
	"github.com/mattslocum/goserver/internal/middleware"
	hash "github.com/mattslocum/goserver/routes/hash"
	shutdown "github.com/mattslocum/goserver/routes/shutdown"
	stats "github.com/mattslocum/goserver/routes/stats"
	"log"
	"net/http"
	"sync"
	"time"
)

var logger = middleware.HttpLoggingMiddleware
var timing = middleware.HttpTimingMiddleware

func setupRoutes(router *http.ServeMux, group *sync.WaitGroup) {
	// TODO: Error handler and logger
	// better pattern matching?
	router.Handle("/hash", logger(timing("GetHash", hash.GetHashRouter(group))))
	router.Handle("/hash/", logger(hash.GetHashRouter(group)))
	router.Handle("/shutdown", logger(new(shutdown.ShutdownRouter)))
	router.Handle("/stats", logger(new(stats.StatsRouter)))
}

func Serve(ctx context.Context) (err error) {
	var group sync.WaitGroup
	router := http.NewServeMux()
	setupRoutes(router, &group)

	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func() {
		log.Print("Starting Server on port 8080")
		// Do we need to do http.Server?
		err = srv.ListenAndServe()
	}()

	// Do we have concurrency issue here?
	if err != nil {
		return err
	}

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer func() {
		cancel()
	}()
	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed: %s", err)
	}

	group.Wait()

	log.Printf("Server stopped.")

	return
}

