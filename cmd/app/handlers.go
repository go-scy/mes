package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
)

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is a middleware")
		next.ServeHTTP(w, r)
	})
}

func employeeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CSRF Token header for HTTP response
		w.Header().Set("X-Csrf-Token", csrf.Token(r))
		fmt.Fprint(w, "employee handler")
	})
}

func productHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "employee handler")
	})
}

func rootHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	})
}
