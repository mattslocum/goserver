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
	"sync"
	"time"
)

var loggerMM = middleware.HttpLoggingMiddleware
var timing = middleware.HttpTimingMiddleware

func setupRoutes(router *http.ServeMux, group *sync.WaitGroup) {
	// TODO: Error handler and logger
	// better pattern matching?
	router.Handle("/hash", loggerMM(timing("GetHash", hash.GetHashRouter(group))))
	router.Handle("/hash/", loggerMM(hash.GetHashRouter(group)))
	router.Handle("/shutdown", loggerMM(new(shutdownRouter.ShutdownRouter)))
	router.Handle("/stats", loggerMM(new(stats.StatsRouter)))
}

func Serve() (err error) {
	var group sync.WaitGroup
	router := http.NewServeMux()
	setupRoutes(router, &group)

	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func() {
		logger.Info.Print("Starting Server on port 8080")
		// Do we need to do http.Server?
		err = srv.ListenAndServe()
	}()

	// Do we have concurrency issue here?
	if err != nil {
		return err
	}

	<-shutdown.Ctx.Done()
	logger.Info.Println("Stopping...")

	ctxTimer, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer func() {
		cancel()
	}()
	if err = srv.Shutdown(ctxTimer); err != nil {
		logger.Error.Fatalf("server Shutdown Failed: %s", err)
	}

	group.Wait()

	logger.Info.Printf("Server stopped.")

	return
}

