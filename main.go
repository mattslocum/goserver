package main

import (
	"github.com/mattslocum/goserver/internal/shutdown"
	"github.com/mattslocum/goserver/routes"
	"log"
	"os"
	"os/signal"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		oscall := <-sig
		log.Printf("System signal: %+v", oscall)
		shutdown.Shutdown()
	}()

	if err := routes.Serve(); err != nil {
		log.Printf("Server Failed. %v\n", err)
	}
}
