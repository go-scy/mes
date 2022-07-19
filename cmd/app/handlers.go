package app

import (
	"fmt"
	"github.com/gorilla/csrf"
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
		err := renderTemplates(w, config, "index.gohtml")
		if err != nil {
			http.Error(w, "cannot render template", http.StatusInternalServerError)
			return
		}
	})
}

// Render files that are stored under "/templates" directory.
func renderTemplates(w http.ResponseWriter, config *AppConfig, html string) error {
	parsedTemplate := config.GetParsedTemplates("templates/" + html)
	return parsedTemplate.Execute(w, nil)
}
