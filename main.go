package main

import (
	"github.com/mattslocum/goserver/internal/logger"
	"github.com/mattslocum/goserver/internal/shutdown"
	"github.com/mattslocum/goserver/routes"
	"os"
	"os/signal"
)

func main() {
	logger.SetupLogs()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		oscall := <-sig
		logger.Info.Printf("System signal: %+v", oscall)
		shutdown.Shutdown()
	}()

	if err := routes.Serve(); err != nil {
		logger.Error.Printf("Server Failed. %v\n", err)
	}
}
