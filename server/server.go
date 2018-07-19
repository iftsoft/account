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
	netSpec := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Printf("Starting, HTTP on: %s\n", netSpec)
	// Create network listener
	listener, err := net.Listen("tcp", netSpec)
	if err != nil {
		fmt.Printf("Error creating listener: %v\n", err)
		return err
	}
	// Define server limits
	server := &http.Server{
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(1) * time.Second,
		MaxHeaderBytes: 1 << 16,
	}
	// Run server asynchronously
	go server.Serve(listener)
	// Wait while server is working
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	fmt.Printf("Got signal: %v, exiting.\n", s)
	// Finish
	return nil
}



