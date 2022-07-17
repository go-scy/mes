package main

import (
	"fmt"
	"net/http"

	"github.com/go-scy/mes/cmd/app"
)

const portNumber = ":8080"

func main() {
	server := app.NewServer()

	fmt.Println("start server listening on ", portNumber)
	http.ListenAndServe(portNumber, server)
}
