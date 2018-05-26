package main

import (
	"net/http"
	"account/handler"
	"account/server"
	"fmt"
)

func main() {
	fmt.Println("-------BEGIN----------")
	// Register default HTTP handlers
	//http.Handle("/", GuiHandler())
	//http.Handle("/gui/", GuiHandler())
	//http.Handle("/static/", ImgHandler())
	//http.Handle("/favicon.ico", ImgHandler())
	http.Handle("/api/", handler.ApiHandler())
	// Run HTTP(S) listener
	server.Run(8080)

	fmt.Println("-------END------------")
}


