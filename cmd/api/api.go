package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/markpassawat/lmwn-assignment/cmd/api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.StandardLogger()

	var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())

	s := server.New()
	// Start server.
	log.Info("starting server...")

	go func() {
		log.Error(s.Run(fmt.Sprintf(":%d", 8080)))
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan

	// cancel()
	wg.Wait()
}
