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
		w.Header().Set("X-Csrf-Token", csrf.Token(r))
		fmt.Fprint(w, "employee handler")
	})
}

func rootHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	})
}
