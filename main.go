package main

import (
	"context"
	"github.com/mattslocum/goserver/routes"
	"log"
	"os"
	"os/signal"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-sig
		log.Printf("System signal: %+v\nStopping...", oscall)
		cancel()
	}()

	if err := routes.Serve(ctx); err != nil {
		log.Printf("Server Failed. %v\n", err)
	}
}
