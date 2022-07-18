package main

import (
	"fmt"
	"net/http"

	"github.com/go-scy/mes/cmd/app"
)

const portNumber = ":8080"

func main() {
	newServer := app.NewServer()

	// Register all handlers
	newServer.AddRoutes()
	newServer.AddMiddleWares()
	//newServer.AddCsrfMiddleware()

	fmt.Println("start newServer listening on ", portNumber)

	http.ListenAndServe(portNumber, newServer)

}
