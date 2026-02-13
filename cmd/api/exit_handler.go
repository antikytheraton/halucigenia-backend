package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// exitHandler blocks until a termination signal is received, allowing the server to gracefully shut down and clean up resources before exiting.
func exitHandler(shutdown func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		log.Printf("Received signal: %s", sig)
		terminate <- true
	}()
	<-terminate

	log.Println("Shutting down server...")
	shutdown()
}
