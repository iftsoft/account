package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func Run(port int) error {
	netSpec := fmt.Sprintf("localhost:%d", port)
	fmt.Println("Starting, HTTP on: %s", netSpec)

	listener, err := net.Listen("tcp", netSpec)
	if err != nil {
		fmt.Println("Error creating listener: %v", err)
		return err
	}
	server := &http.Server{
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(1) * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	go server.Serve(listener)

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	fmt.Println("Got signal: %v, exiting.", s)
}


