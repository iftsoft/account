package main

import (
	"account/front"
	"account/handler"
	"account/server"
	"fmt"
	"net/http"
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
