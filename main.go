package main

import (
	"net/http"
	"account/handler"
	"account/server"
	"fmt"
)

func main() {
	fmt.Println("-------BEGIN----------")
	RunService()
	fmt.Println("-------END------------")
}


func RunService() error {
	// Register default HTTP handler
	//http.Handle("/", GuiHandler())
	//http.Handle("/gui/", GuiHandler())
	//http.Handle("/static/", ImgHandler())
	//http.Handle("/favicon.ico", ImgHandler())
	http.Handle("/api/", handler.ApiHandler())
	// Run HTTP(S) listener
	return server.Run(8080)
}


