package routes

import (
	"context"
	"github.com/mattslocum/goserver/internal/logger"
	"github.com/mattslocum/goserver/internal/middleware"
	"github.com/mattslocum/goserver/internal/shutdown"
	hash "github.com/mattslocum/goserver/routes/hash"
	shutdownRouter "github.com/mattslocum/goserver/routes/shutdown"
	stats "github.com/mattslocum/goserver/routes/stats"
	"net/http"
	"os"
	"sync"
	"time"
)

var loggerMM = middleware.HttpLoggingMiddleware
var timing = middleware.HttpTimingMiddleware

func setupRoutes(router *http.ServeMux, group *sync.WaitGroup) {
	// Could we have better pattern matching? This works for now.
	router.Handle("/hash", loggerMM(timing("PostHash", hash.GetHashRouter(group))))
	router.Handle("/hash/", loggerMM(timing("PostGet", hash.GetHashRouter(group))))
	router.Handle("/shutdown", loggerMM(timing("Shutdown", new(shutdownRouter.ShutdownRouter))))
	router.Handle("/stats", loggerMM(timing("GetStats", new(stats.StatsRouter))))
}

func Serve() (err error) {
	var group sync.WaitGroup
	router := http.NewServeMux()
	setupRoutes(router, &group)

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	srv := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	go func() {
		logger.Info.Printf("Starting Server on port %s", port)
		if err = srv.ListenAndServe(); err != nil {
			logger.Error.Fatalf("Error starting server: %s", err)
		}
	}()

	<-shutdown.Ctx.Done()
	logger.Info.Println("Stopping...")

	ctxTimer, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer func() {
		cancel()
	}()
	if err = srv.Shutdown(ctxTimer); err != nil {
		logger.Error.Printf("Server Shutdown Failed: %s", err)
	}

	group.Wait()

	logger.Info.Printf("Server stopped.")

	return
}

