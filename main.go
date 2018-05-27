package main

import (
	"net/http"
	"account/handler"
	"account/server"
	"account/front"
	"fmt"
)

func main() {
	fmt.Println("-------BEGIN----------")
	// Register default HTTP handlers
	http.Handle("/", front.GuiHandler())
	http.Handle("/api/", handler.ApiHandler())
	// Run HTTP(S) listener
	server.Run(8080)

	fmt.Println("-------END------------")
}


