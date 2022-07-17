package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const portNumber = ":8080"

func main() {
	r := mux.NewRouter()
	r.Handle("/", rootHandler())

	fmt.Print("start server listening on port", portNumber)

	http.ListenAndServe(portNumber, r)
}

func rootHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	})
}
