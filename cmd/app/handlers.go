package app

import (
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"net/http"
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

	})
}

func productHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "employee handler")
	})
}

func rootHandler(config *AppConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedTemplate := template.Must(template.ParseFS(config.EmbeddedFiles, "templates/*"))
		err := parsedTemplate.ExecuteTemplate(w, "index.gohtml", nil)
		if err != nil {
			http.Error(w, "cannot render template", http.StatusInternalServerError)
		}
	})
}
