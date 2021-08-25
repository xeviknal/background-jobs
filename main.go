package main

import (
	"github.com/xeviknal/background-jobs/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	server := server.NewServer()
	server.Start()
	waitForTermination()
	server.Stop()
}

// Waiting until the process receive a Termination signal
func waitForTermination() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	log.Println("Termination Signal Received")
}
