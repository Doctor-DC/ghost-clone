package main

import (
	"github.com/cyantarek/ghost-clone-project/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	// instantiate new server
	s, err := server.New()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	// graceful shutdown setup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		err := s.Run()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	<- stop
	log.Println("Shutting down...")
	err = s.Stop()
	if err != nil {
		log.Fatalf("failed to stop the server: %v", err)
	}
}