package main

import (
	"embed"
	"fmt"
	"github.com/go-scy/mes/cmd/app"
	"net/http"
)

const portNumber = ":8080"

var (
	//go:embed templates
	staticFiles embed.FS
)

func main() {
	appConfig := &app.AppConfig{EmbeddedFiles: staticFiles}
	newServer := app.NewServer(appConfig)

	// Register all handlers
	newServer.AddRoutes()
	newServer.AddMiddleWares()
	newServer.AddCsrfMiddleware()

	fmt.Println("start server listening on ", portNumber)

	http.ListenAndServe(portNumber, newServer.Router())

}
